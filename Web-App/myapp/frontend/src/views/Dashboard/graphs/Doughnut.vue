<script>
import { Doughnut, mixins } from 'vue-chartjs'

export default {
  name: 'DoughnutChart',

  extends: Doughnut,
  mixins: [mixins.reactiveProp],
  props: {
    options: {
      type: Object,
      default: () => {
        return {}
      }
    }
  },
  mounted() {
    this.renderChart(this.chartData, this.options)
  }
}
</script>
