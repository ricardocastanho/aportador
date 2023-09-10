<template>
  <v-container class="fill-height">
    <v-responsive class="align-center text-center fill-height mx-auto" max-width="1000">
      <h1 class="text-h2 font-weight-bold mb-2">Aportador</h1>
      <div class="text-body-2 font-weight-light mb-6">
        Descubra o preço teto das suas ações
      </div>
      <v-text-field
        v-model="ticker"
        label="Ticker"
        variant="solo"
        class="ma-6"
        append-icon="mdi-magnify"
        @click:append="search"
      />

      <div class="my-5"></div>

      <v-table>
        <thead>
          <tr>
            <th class="text-left">
              Ticker
            </th>
            <th class="text-left">
              Preço atual
            </th>
            <th class="text-left">
              Preço teto Grahan
            </th>
            <th class="text-left">
              Preço teto Bazin
            </th>
            <th class="text-left">
              Margem de segurança Grahan
            </th>
            <th class="text-left">
              Margem de segurança Bazin
            </th>
          </tr>
        </thead>
        <tbody>
          <tr
            v-for="stock in stocks"
            :key="stock.ticker"
          >
            <td>{{ stock.ticker }}</td>
            <td>R$ {{ stock.actualPrice.toFixed(2) }}</td>
            <td>R${{ stock.grahanFairPrice.toFixed(2) }}</td>
            <td>R${{ stock.bazinFairPrice.toFixed(2) }}</td>
            <td>{{ stock.grahanSafeMargin.toFixed(0) }}%</td>
            <td>{{ stock.bazinSafeMargin.toFixed(0) }}%</td>
          </tr>
        </tbody>
      </v-table>
    </v-responsive>
  </v-container>
</template>

<script lang="ts" setup>
import { ref, onBeforeMount } from "vue"

type Stock = {
  ticker: string
  actualPrice: number
  grahanFairPrice: number
  bazinFairPrice: number
  grahanSafeMargin: number
  bazinSafeMargin: number
}

type Response = {
  data: Stock[]
  error: string
}

const DEFAULT_STOCKS: string[] = ["BBAS3", "TAEE11", "BBSE3"]
const API_URL = import.meta.env.VITE_API_URL || ""

const tickers = ref<string[]>(DEFAULT_STOCKS)
const ticker = ref<string>("")
const stocks = ref<Stock[]>([])

const search = async (): Promise<void> => {
  tickers.value.push(ticker.value)
  await getStocks([ticker.value])
}

const getStocks = async (tickers: string[]): Promise<void> => {
  let url = API_URL

  url = url.concat("/search?")
  tickers.forEach((ticker, i) => {
    if (i !== 0) {
      url = url.concat("&")
    }
    url = url.concat(`stock=${ticker}`)
  })

  const res = await fetch(url)
  const dataText = await res.text()
  const { data }: Response = JSON.parse(dataText)
  stocks.value.push(...data)
}

onBeforeMount(async () => {
  await getStocks(DEFAULT_STOCKS)
})

</script>
