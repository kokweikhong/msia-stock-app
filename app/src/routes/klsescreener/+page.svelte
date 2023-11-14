<script lang="ts">
  import { onMount } from "svelte";
  import { GenerateStockAreaChart } from "$lib";
  import Favourite from "../../components/Favourite.svelte";

  type StockData = {
    date: string;
    open: number;
    close: number;
    high: number;
    low: number;
    volume: number;
  };
  // generate fake data for stock area chart, data includes date, open, close, high, low, volume
  const data: StockData[] = [
    {
      date: "2020-01-01",
      open: 100,
      close: 200,
      high: 300,
      low: 50,
      volume: 1000,
    },
    {
      date: "2020-01-02",
      open: 200,
      close: 300,
      high: 400,
      low: 150,
      volume: 2000,
    },
    {
      date: "2020-01-03",
      open: 300,
      close: 400,
      high: 500,
      low: 250,
      volume: 3000,
    },
    {
      date: "2020-01-04",
      open: 400,
      close: 500,
      high: 600,
      low: 350,
      volume: 4000,
    },
    {
      date: "2020-01-05",
      open: 500,
      close: 600,
      high: 700,
      low: 450,
      volume: 5000,
    },
    {
      date: "2020-01-06",
      open: 600,
      close: 700,
      high: 800,
      low: 550,
      volume: 6000,
    },
    {
      date: "2020-01-07",
      open: 700,
      close: 800,
      high: 900,
      low: 650,
      volume: 7000,
    },
    {
      date: "2020-01-08",
      open: 800,
      close: 900,
      high: 1000,
      low: 750,
      volume: 8000,
    },
    {
      date: "2020-01-09",
      open: 900,
      close: 1000,
      high: 1100,
      low: 850,
      volume: 9000,
    },
  ];

  onMount(() => {
    GenerateStockAreaChart("main", data);
    GenerateStockAreaChart("stock-chart", data);
  });
</script>

<div class="flex min-h-full flex-col">
  <div
    class="mx-auto flex w-full max-w-7xl items-start gap-x-8 px-4 py-10 sm:px-6 lg:px-8"
  >
    <aside class="sticky top-8 w-44 shrink-0 lg:block">
      <div>
        <ul role="list" class="divide-y divide-gray-100">
          <li class="flex flex-wrap items-center justify-between gap-x-6 py-5">
            <div class="flex gap-x-4">
              <div class="min-w-0 flex-auto">
                <p class="text-sm font-semibold leading-6 text-gray-900">
                  Leslie Alexander
                </p>
                <p class="mt-1 truncate text-xs leading-5 text-gray-500">
                  leslie.alexander@example.com
                </p>
              </div>
            </div>
            <div class="py-1">
              <Favourite isFavourite={false} />
            </div>
          </li>
          <li class="flex flex-wrap items-center justify-between gap-x-6 py-5">
            <div class="flex gap-x-4">
              <div class="min-w-0 flex-auto">
                <p class="text-sm font-semibold leading-6 text-gray-900">
                  Leslie Alexander
                </p>
                <p class="mt-1 truncate text-xs leading-5 text-gray-500">
                  leslie.alexander@example.com
                </p>
              </div>
            </div>
            <div class="py-1">
              <Favourite isFavourite={false} />
            </div>
          </li>
          <li class="flex flex-wrap items-center justify-between gap-x-6 py-5">
            <div class="flex gap-x-4">
              <div class="min-w-0 flex-auto">
                <p class="text-sm font-semibold leading-6 text-gray-900">
                  Leslie Alexander
                </p>
                <p class="mt-1 truncate text-xs leading-5 text-gray-500">
                  leslie.alexander@example.com
                </p>
              </div>
            </div>
            <div class="py-1">
              <Favourite isFavourite={false} />
            </div>
          </li>
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
              <h2>Stock Area Chart</h2>
              <!-- Content goes here -->
              <!-- We use less vertical padding on card headers on desktop than on body sections -->
            </div>
            <div class="px-4 py-5 sm:p-6">
              <div id="stock-chart" class="w-full h-[350px]" />
            </div>
          </div>
        </li>
        <li class="py-4">
          <div
            class="divide-y divide-gray-200 overflow-hidden rounded-lg bg-white shadow"
          >
            <div class="px-4 py-5 sm:px-6">
              <h2>Stock Area Chart</h2>
              <!-- Content goes here -->
              <!-- We use less vertical padding on card headers on desktop than on body sections -->
            </div>
            <div class="px-4 py-5 sm:p-6">
              <div id="main" class="w-full h-[350px]" />
            </div>
          </div>
        </li>
      </ul>
    </main>
  </div>
</div>
