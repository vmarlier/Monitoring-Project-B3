<script>
import { Bar, mixins } from 'vue-chartjs'
// BAR
export default {
  name: 'BarChart',
  extends: Bar,
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
