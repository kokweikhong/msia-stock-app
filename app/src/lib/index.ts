import * as echarts from "echarts";
import type { OHLC } from "../types/klsescreener/klsescreener";

export function GenerateStockAreaChart(divID: string, data: OHLC[]) {
  const chartDom = document.getElementById(divID);
  if (!chartDom) return;
  if (echarts.getInstanceByDom(chartDom) != null) {
    echarts.getInstanceByDom(chartDom)?.dispose();
  }
  const myChart = echarts.init(chartDom);
  // let option: echarts.EChartsOption;
  const option: echarts.EChartsOption = {
    tooltip: {
      trigger: "axis",
      position: function (pt) {
        return [pt[0], "10%"];
      },
    },
    title: {
      left: "center",
      // text: "Stock Area Chart",
    },
    toolbox: {
      feature: {
        dataZoom: {
          yAxisIndex: "none",
        },
        restore: {},
        saveAsImage: {},
      },
    },
    xAxis: {
      type: "category",
      boundaryGap: false,
      data: data?.map((x: OHLC) => x.date),
    },
    yAxis: {
      type: "value",
      boundaryGap: [0, "100%"],
      splitLine: {
        show: false,
      },
      min(extent) {
        return extent.min - 0.5;
      },
      max(extent) {
        return extent.max + 0.5;
      },
    },
    dataZoom: [
      {
        type: "inside",
        start: 0,
        end: 100,
      },
      {
        start: 0,
        end: 100,
      },
    ],
    series: [
      {
        name: "Stock Area Chart",
        type: "line",
        smooth: true,
        symbol: "none",
        sampling: "average",
        itemStyle: {
          color: "rgb(255, 70, 131)",
        },
        areaStyle: {
          color: new echarts.graphic.LinearGradient(0, 0, 0, 1, [
            {
              offset: 0,
              color: "rgb(255, 158, 68)",
            },
            {
              offset: 1,
              color: "rgb(255, 70, 131)",
            },
          ]),
        },
        data: data?.map((x: OHLC) => x.close),
      },
    ],
  };

  option && myChart.setOption(option);
}
