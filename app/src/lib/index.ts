// place files you want to import through the `$lib` alias in this folder.
import * as echarts from "echarts";
export function GenerateStockAreaChart(divID: string, data: any) {
  var chartDom = document.getElementById(divID);
  if (!chartDom) return;
  if (echarts.getInstanceByDom(chartDom) != null) {
    echarts.getInstanceByDom(chartDom)?.dispose();
  }
  var myChart = echarts.init(chartDom);
  var option: echarts.EChartsOption;
  option = {
    tooltip: {
      trigger: "axis",
      position: function (pt) {
        return [pt[0], "10%"];
      },
    },
    title: {
      left: "center",
      text: "Stock Area Chart",
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
      data: data?.map((x: any) => x.date),
    },
    yAxis: {
      type: "value",
      boundaryGap: [0, "100%"],
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
        data: data?.map((x: any) => x.close),
      },
    ],
  };

  option && myChart.setOption(option);
}
