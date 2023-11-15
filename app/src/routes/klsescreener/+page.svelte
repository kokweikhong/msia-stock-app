<script lang="ts">
  import { onMount } from "svelte";
  import { GenerateStockAreaChart } from "$lib";
  import Favourite from "../../components/Favourite.svelte";
  import {
    getKLSEScreenerQuotes,
    getKLSEScreenerIndexData,
    getKLSEScreenerStockData,
  } from "$lib/klsescreener";
  import type { OHLC, Quote } from "../../types/klsescreener/klsescreener";

  let quotes = getKLSEScreenerQuotes() as unknown as Quote[];
  let indexData: OHLC[];
  let stockData: OHLC[];
  let selectedQuote: Quote;

  async function getStockData(code: string) {
    stockData = await getKLSEScreenerStockData(code);
    // stockData based on quotes first date
    stockData = stockData.filter((data) => {
      return data.date >= indexData[0].date;
    });
    GenerateStockAreaChart("stock-chart", stockData);
  }

  onMount(async () => {
    indexData = await getKLSEScreenerIndexData();
    GenerateStockAreaChart("main", indexData);
  });
</script>

<div class="flex min-h-full flex-col">
  <div
    class="mx-auto flex w-full max-w-7xl items-start gap-x-8 px-4 py-10 sm:px-6 lg:px-8"
  >
    <aside class="sticky top-8 w-48 shrink-0 lg:block">
      <div class="max-h-[700px] overflow-y-auto w-full">
        <ul role="list" class="divide-y divide-gray-100">
          {#await quotes}
            Loading planet...
          {:then quotes}
            <!-- Your next planet is {quotes}. -->
            {#each quotes as quote}
              <li
                class="flex flex-col items-start justify-between gap-x-6 py-5 {selectedQuote?.code ===
                quote.code
                  ? 'bg-gray-100'
                  : ''}}"
              >
                <div class="flex gap-x-4">
                  <div class="min-w-0 flex-auto">
                    <p class="text-sm font-semibold leading-6 text-gray-900">
                      {quote.name}
                    </p>
                    <p class="mt-1 truncate text-xs leading-5 text-gray-500">
                      {quote.code}
                    </p>
                  </div>
                </div>
                <div class="py-1">
                  <Favourite isFavourite={false} />
                  <button
                    on:click={() => {
                      getStockData(quote.code);
                      selectedQuote = quote;
                    }}
                  >
                    <svg
                      xmlns="http://www.w3.org/2000/svg"
                      width="24"
                      height="24"
                      viewBox="0 0 24 24"
                      fill="none"
                      stroke="currentColor"
                      stroke-width="2"
                      stroke-linecap="round"
                      stroke-linejoin="round"
                      class="lucide lucide-line-chart"
                      ><path d="M3 3v18h18" /><path
                        d="m19 9-5 5-4-4-3 3"
                      /></svg
                    >
                  </button>
                </div>
              </li>
            {/each}
          {:catch someError}
            System error: {someError.message}.
          {/await}
        </ul>
      </div>
    </aside>

    <main class="flex-1">
      <ul role="list" class="divide-y divide-gray-200">
        <li class="py-4">
          <div
            class="divide-y divide-gray-200 overflow-hidden rounded-lg bg-white shadow"
          >
            <div class="px-4 py-5 sm:px-6">
              <h2>KLSE Index</h2>
              <!-- Content goes here -->
              <!-- We use less vertical padding on card headers on desktop than on body sections -->
            </div>
            <div class="px-4 py-5 sm:p-6">
              <div id="main" class="w-full h-[350px]" />
            </div>
          </div>
        </li>
        <li class="py-4 {stockData ? 'opacity-100' : 'opacity-0'}">
          <div
            class="divide-y divide-gray-200 overflow-hidden rounded-lg bg-white shadow"
          >
            <div class="px-4 py-5 sm:px-6">
              {#await quotes}
                <span>Loading...</span>
              {:then quotes}
                <div class="flex justify-between">
                  <h2>
                    {selectedQuote
                      ? selectedQuote?.code + " | " + selectedQuote?.name
                      : ""}
                  </h2>

                  <span
                    class="isolate inline-flex items-center rounded-md shadow-sm"
                  >
                    <button
                      type="button"
                      class="relative inline-flex items-center rounded-l-md bg-white px-2 py-2 text-gray-400 ring-1 ring-inset ring-gray-300 hover:bg-gray-50 focus:z-10"
                      disabled={selectedQuote
                        ? selectedQuote?.code === quotes[0]?.code
                        : false}
                      on:click={async () => {
                        let index = quotes.findIndex(
                          (quote) => quote.code === selectedQuote?.code
                        );
                        await getStockData(quotes[index - 1].code);
                        selectedQuote = quotes[index - 1];
                      }}
                    >
                      <span class="sr-only">Previous</span>
                      <svg
                        class="h-5 w-5"
                        viewBox="0 0 20 20"
                        fill="currentColor"
                        aria-hidden="true"
                      >
                        <path
                          fill-rule="evenodd"
                          d="M12.79 5.23a.75.75 0 01-.02 1.06L8.832 10l3.938 3.71a.75.75 0 11-1.04 1.08l-4.5-4.25a.75.75 0 010-1.08l4.5-4.25a.75.75 0 011.06.02z"
                          clip-rule="evenodd"
                        />
                      </svg>
                    </button>
                    <button
                      type="button"
                      class="relative -ml-px inline-flex items-center rounded-r-md bg-white px-2 py-2 text-gray-400 ring-1 ring-inset ring-gray-300 hover:bg-gray-50 focus:z-10"
                      disabled={selectedQuote
                        ? selectedQuote?.code ===
                          quotes[quotes.length - 1]?.code
                        : false}
                      on:click={async () => {
                        if (quotes.length === 0) return;
                        let index = quotes.findIndex(
                          (quote) => quote.code === selectedQuote?.code
                        );
                        await getStockData(quotes[index + 1].code);
                        selectedQuote = quotes[index + 1];
                      }}
                    >
                      <span class="sr-only">Next</span>
                      <svg
                        class="h-5 w-5"
                        viewBox="0 0 20 20"
                        fill="currentColor"
                        aria-hidden="true"
                      >
                        <path
                          fill-rule="evenodd"
                          d="M7.21 14.77a.75.75 0 01.02-1.06L11.168 10 7.23 6.29a.75.75 0 111.04-1.08l4.5 4.25a.75.75 0 010 1.08l-4.5 4.25a.75.75 0 01-1.06-.02z"
                          clip-rule="evenodd"
                        />
                      </svg>
                    </button>
                    <span class="p-2 ml-4">
                      {quotes.findIndex(
                        (quote) => quote.code === selectedQuote?.code
                      ) + 1}
                      /{quotes.length}
                    </span>
                  </span>
                </div>
              {/await}
              <!-- Content goes here -->
              <!-- We use less vertical padding on card headers on desktop than on body sections -->
            </div>
            <div class="px-4 py-5 sm:p-6">
              <div id="stock-chart" class="w-full h-[350px]" />
            </div>
          </div>
        </li>
      </ul>
    </main>
  </div>
</div>
