<template>
  <v-container class="px-16">
    <v-row class="align-center text-center mt-16">
      <v-col>
        <h1 class="text-h2 font-weight-bold mb-2">Aportador</h1>
        <div class="text-body-2 font-weight-light mb-6">
          Descubra o preço teto das suas ações
        </div>
        <v-text-field
          :value="ticker"
          label="Ticker"
          variant="solo"
          class="ma-6"
          append-icon="mdi-magnify"
          @update:model-value="(t) => ticker = toUpperCase(t)"
          @click:append="search"
        />
      </v-col>
    </v-row>

    <v-row class="align-center text-center">
      <v-col>
        <div class="my-5"></div>

        <v-table>
          <thead>
            <tr>
              <th class="text-left">
                Ticker
              </th>
              <th class="text-left">
                Preço
              </th>
              <th class="text-left">
                Yield Mínimo
              </th>
              <th class="text-left">
                Payout
              </th>
              <th class="text-left">
                Lucro Projetado
              </th>
              <th class="text-left">
                DPA
              </th>
              <th class="text-left">
                Preço teto
              </th>
              <th class="text-left"/>
            </tr>
          </thead>
          <tbody>
            <tr
              v-for="(stock, i) in stocks"
              :key="stock.ticker"
            >
              <td>{{ stock.ticker }}</td>
              <td>R$ {{ stock.price }}</td>
              <td>
                <v-text-field suffix="%" v-model.number="input[stock.ticker].dy" @update:model-value="() => persistInput()" />
              </td>
              <td>
                <v-text-field suffix="%" v-model.number="input[stock.ticker].payout" @update:model-value="() => persistInput()" />
              </td>
              <td>
                <v-text-field prefix="R$" v-model.string="input[stock.ticker].profit" @update:model-value="() => persistInput()" />
              </td>
              <td>
                R$ {{ stock.dpa .toFixed(2)}}
              </td>
              <td :class="parseAmount(stock.price) < stock.ceilPrice ? 'text-success' : 'text-error'">
                R$ {{ stock.ceilPrice.toFixed(2) }}
              </td>
              <td>
                <v-icon icon="mdi-close" @click="() => remove(stock.ticker)" />
              </td>
            </tr>
          </tbody>
        </v-table>
      </v-col>
    </v-row>
  </v-container>
</template>

<script lang="ts" setup>
import { ref, onBeforeMount, computed, reactive } from "vue"

import { Stock, calculateDpa, getCeilPrice } from "../models/stock"
import { parseAmount, toUpperCase } from "../utils"

type StockRaw = {
  ticker: string
  price: string
  shares: string
  profit: string
  payout: string
}

type Response = {
  data: StockRaw[]
  error: string
}

type Input = {
  [ticker: string]: {
    dy: number,
    payout: number,
    profit: string,
  }
}

const API_URL = import.meta.env.VITE_API_URL || ""
const savedInput: Input = JSON.parse(localStorage.getItem("data") ?? "{}")

const input = reactive<Input>(savedInput)
const ticker = ref<string>("")
const stocksRaw = ref<StockRaw[]>([])

const stocks = computed<Stock[]>(() => {
  return stocksRaw.value.map((stock: StockRaw): Stock => {
    const dpa = calculateDpa(input[stock.ticker].profit, stock.shares, input[stock.ticker].payout)

    return {
      ticker: stock.ticker,
      price: stock.price,
      shares: stock.shares,
      profit: stock.profit,
      payout: input[stock.ticker].payout.toString(),
      dpa: dpa,
      ceilPrice: getCeilPrice(dpa, input[stock.ticker].dy)
    }
  })
})

const search = async (): Promise<void> => {
  try {
    await getStocks([...Object.keys(input), ticker.value])
    ticker.value = ""
  } catch (err) {
    console.error("Error searching stocks:", err);
  }
}

const remove = async (ticker: string) => {
  const index = stocksRaw.value.findIndex((stock: StockRaw): boolean => stock.ticker === ticker)
  stocksRaw.value.splice(index, 1)
  delete input[ticker]
  localStorage.setItem("data", JSON.stringify(input))
}

const persistInput = () => {
  localStorage.setItem("data", JSON.stringify(input))
}

const getStocks = async (tickers: string[]): Promise<void> => {
  if (tickers.length === 0) {
    return
  }

  let url = API_URL

  url = url.concat("/stocks?")

  tickers.forEach((ticker, i) => {
    if (i !== 0) {
      url = url.concat("&")
    }
    url = url.concat(`stock=${ticker}`)
  })

  const res = await fetch(url)
  const dataText = await res.text()
  const { data }: Response = JSON.parse(dataText)

  stocksRaw.value = [...data]

  stocksRaw.value.forEach((s: StockRaw): void => {
    input[s.ticker] = {
      dy: input[s.ticker]?.dy ?? 6,
      payout: input[s.ticker]?.payout ?? 30,
      profit: input[s.ticker]?.profit ?? s.profit,
    }
  }, {} as Input)

  localStorage.setItem("data", JSON.stringify(input))
}

onBeforeMount(async () => {
  await getStocks(Object.keys(input))
})
</script>
