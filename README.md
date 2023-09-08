# aportador

Find Brazillian stocks best prices following Grahan and Bazin principles.
Currently scraping [Fundamentus](https://www.fundamentus.com.br) data.

## Motivation

Bazin's long-term investment approach provides a remedy for the common pitfall of impulsive trading. Novice investors often fall victim to emotional decision-making, leading to buying and selling based on short-term market fluctuations. Bazin's emphasis on patience and holding stocks for extended periods encourages beginners to avoid the trap of overtrading and instead focus on the fundamental value of companies. By promoting a disciplined, steady approach, Bazin's principles help new investors overcome the challenge of emotional trading.

Benjamin Graham's value investing philosophy addresses the difficulty of identifying undervalued opportunities in a sea of market noise. For beginner investors, distinguishing between overhyped stocks and genuinely undervalued assets can be daunting. Graham's emphasis on analyzing a company's intrinsic value, considering its financials and market position, provides a systematic approach to make informed investment decisions. By encouraging a thorough assessment of a company's fundamentals, Graham's principles guide novice investors in overcoming the challenge of identifying solid investment prospects amidst market fluctuations.

In essence, the principles of Bazin and Graham provide valuable guidance for novice investors by tackling the challenges of emotional decision-making and identifying undervalued assets. Their time-tested approaches offer clarity, discipline, and a systematic methodology, ultimately helping beginners build a strong foundation for successful investing in the dynamic world of finance.

## Who is Bazin?

Luiz Bazin is a renowned Brazilian investor known for his long-term approach and investment in dividend-paying companies. He began his career in the 1960s and built his fortune through the purchase of shares in solid and well-established companies, mainly focusing on sectors such as banking and energy. Bazin is notable for consistently reinvesting the dividends received, accumulating a robust portfolio over the decades. His success story has inspired many investors to adopt long-term investment strategies and to appreciate the importance of dividends.

## Who is Graham?

Warren Graham (known as "Graham") is a fictional investor created by Benjamin Graham, considered the "father" of value investing. Graham is notable for his disciplined, fundamental analysis-based approach. His classic book "The Intelligent Investor" influenced generations of investors throughout the 20th century and remains a reference in the field of investments. Through his ideas, Graham promoted the importance of buying shares of undervalued companies relative to their intrinsic value, seeking margins of safety and minimizing risks.

## Commands

Searching Brazilian stocks fair values. You can find the tickers [here](https://www.b3.com.br/pt_br/produtos-e-servicos/negociacao/renda-variavel/empresas-listadas.htm).

```shell
go run main.go search --tickers BBAS3,TAEE11,PETR4
```

Searching Brazilian stocks fair prices to get 10 percent of dividend yield considering the last 2 years of dividend history.

```shell
go run main.go search --tickers BBAS3,TAEE11,PETR4 --dividend-yield 10 --dividend-years 2
```
