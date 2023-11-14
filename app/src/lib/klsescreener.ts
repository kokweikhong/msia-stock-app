import type { Quote, OHLC } from "../types/klsescreener/klsescreener";

export async function getKLSEScreenerQuotes(): Promise<Quote[]> {
  // sleep(1000);
  await sleep(1000);
  const response = await fetch("http://localhost:8080/klsescreener/quotes");
  const json = await response.json();
  return json;
}

export async function getKLSEScreenerIndexData(): Promise<OHLC[]> {
  await sleep(1000);
  const response = await fetch("http://localhost:8080/klsescreener/index");
  const json = await response.json();
  return json;
}

export async function getKLSEScreenerStockData(code: string): Promise<OHLC[]> {
  await sleep(1000);
  const response = await fetch(
    `http://localhost:8080/klsescreener/stock/${code}`
  );
  const json = await response.json();
  return json;
}

function sleep(ms: number) {
  return new Promise((resolve) => setTimeout(resolve, ms));
}
