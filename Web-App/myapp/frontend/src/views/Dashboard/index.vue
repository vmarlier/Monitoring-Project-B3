<template>
  <section class="app-main">
    <transition name="fade" mode="out-in">
      <div class="row">
        <div class="col-12">
          <div class="row row-eq-height">
            <div class="col-md-12 text-left">
              <div class="content m-2">
                <h4 class="text-center">Details</h4>
                <div class="row">
                  <div class="col-6 pl-4 border-right">
                    <div class="row">
                      <div class="col-12">
                        <p class="card-title font-weight-bold">General</p>
                        <div class="row">
                          <div class="col-7">IP:</div>
                          <div class="col-5">{{ serverData.hostname }}</div>
                        </div>
                        <div class="row">
                          <div class="col-7">OS:</div>
                          <div class="col-5">{{ serverData.os }}</div>
                        </div>
                      </div>
                    </div>
                  </div>
                  <div class="col-6 pl-4">
                    <div class="row">
                      <div class="col-12">
                        <p class="card-title font-weight-bold">Running</p>
                        <div class="row">
                          <div class="col-7">Uptime:</div>
                          <div class="col-5">{{ uptime }}</div>
                        </div>
                        <div class="row">
                          <div class="col-7">Processes:</div>
                          <div class="col-5">{{ serverData.process }}</div>
                        </div>
                      </div>
                    </div>
                  </div>
                </div>
              </div>
            </div>
          </div>
          <div class="row row-eq-height">
            <div class="col-md-4 h-100">
              <div class="content m-2">
                <div class="row">
                  <div class="col-6 text-left">
                    <h4>RAM</h4>
                  </div>
                  <div class="col-6 text-right">
                    <h3 id="ramPerc" :style="{ color: assignColorPerc(serverData.memory.usedPerc) }">
                      {{ serverData.memory.usedPerc || 0 }}%
                    </h3>
                  </div>
                </div>
                <doughnut-chart
                  :chart-data="ramData"
                  :options="ramOptions"
                  class="chart"
                  chart-id="ram"
                ></doughnut-chart>
              </div>
            </div>
            <div class="col-md-4 h-100">
              <div class="content m-2">
                <div class="row">
                  <div class="col-4 text-left">
                    <h4>CPU</h4>
                  </div>
                  <div class="col-4">
                    <div v-if="cpuTabs.length && cpuTabs.length > 1" class="btn-group">
                      <button
                        type="button"
                        class="btn btn-outline-dark dropdown-toggle btn-sm"
                        data-toggle="dropdown"
                        aria-haspopup="true"
                        aria-expanded="false"
                      >
                        {{ selectedCpuTab }}
                      </button>
                      <div class="dropdown-menu">
                        <a
                          v-for="(tab, index) in cpuTabs"
                          :key="index"
                          :class="{ active: tab === selectedCpuTab }"
                          class="dropdown-item"
                          @click.prevent="selectedCpuTab = tab"
                        >
                          {{ tab }}
                        </a>
                      </div>
                    </div>
                  </div>
                  <div class="col-4 text-right">
                    <h3 id="cpuPerc" :style="{ color: assignColorPerc(cpuPerc) }">{{ cpuPerc }}%</h3>
                  </div>
                </div>
                <bar-chart :chart-data="selectedCpuData" :options="cpuOptions" class="chart" chart-id="cpu"></bar-chart>
              </div>
            </div>
            <div class="col-md-4 h-100">
              <div class="content m-2">
                <div class="row">
                  <div class="col-6 text-left">
                    <h4>Disks</h4>
                  </div>
                  <div class="col-6 text-right">
                    <h3 id="diskPerc" :style="{ color: assignColorPerc(diskPerc) }">{{ diskPerc }}%</h3>
                  </div>
                  <ul v-if="diskTabs.length && diskTabs.length > 1" class="nav nav-pills nav-justified pl-2 pb-2">
                    <li
                      v-for="(tab, index) in diskTabs"
                      :key="index"
                      :class="{ active: tab === selectedDiskTab }"
                      class="nav-item tabs-item mr-1"
                    >
                      <a class="nav-link" href @click.prevent="selectedDiskTab = tab">{{ tab }}</a>
                    </li>
                  </ul>
                </div>
                <doughnut-chart
                  :chart-data="selectedDiskData"
                  :options="diskOptions"
                  class="chart"
                  chart-id="disk"
                ></doughnut-chart>
              </div>
            </div>
          </div>
        </div>
      </div>
    </transition>
  </section>
</template>

<script>
import BarChart from './graphs/Bar'
import DoughnutChart from './graphs/Doughnut'

import { getMetrics } from '@/api/metrics'

import moment from 'moment'

export default {
  name: 'Dashboard',
  components: { BarChart, DoughnutChart },
  data() {
    return {
      alertStatus: {
        cpu: false,
        disk: false,
        ram: false
      },
      colorsPercs: {
        green: '#11a786',
        yellow: '#edc110',
        orange: '#e29416',
        red: '#e20909'
      },
      cpuOptions: {
        scales: {
          yAxes: [
            {
              display: true,
              ticks: {
                beginAtZero: true,
                stepSize: 5,
                max: 100
              }
            }
          ]
        },
        legend: {
          display: false
        },
        tooltips: {
          callbacks: {
            label: function(tooltipItem) {
              return tooltipItem.yLabel
            }
          }
        }
      },
      diskOptions: {
        elements: {
          arc: {
            borderWidth: 0
          }
        }
      },
      loading: true,
      ramOptions: {
        elements: {
          arc: {
            borderWidth: 0
          }
        }
      },
      selectedCpuData: {
        labels: [],
        datasets: []
      },
      selectedCpuTab: '',

      selectedDiskData: {
        labels: [],
        datasets: []
      },
      selectedDiskTab: '',
      serverData: {
        hostname: '',
        os: '',
        uptime: '',
        process: '',
        memory: {
          total: '',
          free: '',
          usedPerc: ''
        },
        disk: [],
        cpus: [],
        interfaces: []
      }
    }
  },
  computed: {
    cpuData() {
      return this.serverData.cpus.map(cpu => {
        const labels = []
        cpu.threaduse.forEach((thread, index) => {
          labels.push(`thread ${index + 1}`)
        })
        return {
          labels,
          datasets: [
            {
              backgroundColor: '#11a786',
              data: cpu.threaduse.map(thread => {
                if (thread > 0) {
                  return Math.round(thread * 10) / 10
                }
                return thread
              })
            }
          ]
        }
      })
    },
    cpuTabs() {
      return this.serverData.cpus.map((cpu, index) => `CPU ${index + 1}`)
    },
    diskData() {
      return this.serverData.disk.map(disk => {
        return {
          labels: ['Used', 'Free'],
          datasets: [
            {
              label: 'Ram usage',
              backgroundColor: ['#11a786', '#e5e5e5'],
              data: [disk.used, disk.free]
            }
          ]
        }
      })
    },
    diskTabs() {
      return this.serverData.disk.map((disk, index) => `Disk ${index + 1}`)
    },
    ramData() {
      const used = this.serverData.memory.total * this.serverData.memory.usedPerc / 100
      const free = this.serverData.memory.total - used
      const roundUsed = used > 0 ? Math.round(used * 10) / 10 : used
      const roundFree = free > 0 ? Math.round(free * 10) / 10 : free
      return {
        labels: ['Used', 'Free'],
        datasets: [
          {
            label: 'Ram usage',
            backgroundColor: ['#11a786', '#e5e5e5'],
            data: [roundUsed, roundFree]
          }
        ]
      }
    },
    cpuPerc() {
      if (
        this.selectedCpuData &&
        this.selectedCpuData.datasets &&
        this.selectedCpuData.datasets.length &&
        this.selectedCpuData.datasets[0].data &&
        this.selectedCpuData.datasets[0].data.length
      ) {
        const sum = this.selectedCpuData.datasets[0].data.reduce(function(acc, curr) {
          return acc + curr
        })
        return Math.round(sum / this.selectedCpuData.datasets[0].data.length * 10) / 10
      }
      return 0
    },
    diskPerc() {
      const index = this.diskTabs.indexOf(this.selectedDiskTab)
      if (index !== -1) {
        return this.serverData.disk[index].usedPerc
      }
      return 0
    },
    uptime() {
      const duration = moment.duration(parseInt(this.serverData.uptime), 'seconds')
      const days = duration.days()
      const hours = duration.hours()
      const minutes = duration.minutes()
      const seconds = duration.seconds()
      if (duration >= moment.duration(24, 'hours')) {
        return `${days}d ${hours}h ${minutes}min ${seconds}s`
      } else if (duration < moment.duration(60, 'minutes')) {
        return `${minutes} minutes ${seconds} seconds`
      } else if (duration < moment.duration(60, 'seconds')) {
        return `${seconds} secondes`
      } else {
        return `${hours}h ${minutes}min ${seconds}s`
      }
    }
  },
  watch: {
    selectedCpuTab: {
      immediate: true,
      handler() {
        const index = this.cpuTabs.indexOf(this.selectedCpuTab)
        if (index !== -1) {
          this.selectedCpuData = JSON.parse(JSON.stringify(this.cpuData[index]))
        }
      }
    },
    selectedDiskTab: {
      immediate: true,
      handler() {
        const index = this.diskTabs.indexOf(this.selectedDiskTab)
        if (index !== -1) {
          this.selectedDiskData = JSON.parse(JSON.stringify(this.diskData[index]))
        }
      }
    }
  },
  created() {
    this.getMetrics()
  },
  mounted() {
    this.$nextTick(function() {
      this.interval = window.setInterval(() => {
        this.getMetrics()
      }, 1000)
    })
  },
  beforeDestroy() {
    clearInterval(this.interval)
  },
  methods: {
    assignColorPerc(value) {
      return value < 50
        ? this.colorsPercs.green
        : value < 75 ? this.colorsPercs.yellow : value < 90 ? this.colorsPercs.orange : this.colorsPercs.red
    },
    getMetrics() {
      getMetrics().then(response => {
        if (response && response.data) {
          this.serverData = response.data
          if (this.cpuTabs.length && !this.cpuTabs.includes(this.selectedCpuTab)) {
            this.selectedCpuTab = this.cpuTabs[0]
          }
          const indexCpu = this.cpuTabs.indexOf(this.selectedCpuTab)
          if (indexCpu !== -1) {
            this.selectedCpuData = JSON.parse(JSON.stringify(this.cpuData[indexCpu]))
          }

          if (this.diskTabs.length && !this.diskTabs.includes(this.selectedDiskTab)) {
            this.selectedDiskTab = this.diskTabs[0]
          }
          const indexDisk = this.diskTabs.indexOf(this.selectedDiskTab)
          if (indexDisk !== -1) {
            this.selectedDiskData = JSON.parse(JSON.stringify(this.diskData[indexDisk]))
          }
        }
      })
    }
  }
}
</script>

<style lang="scss">
@import '@/style/_variables.scss';

$positiveColor: $main-color;
$mediumColor: rgb(245, 139, 18);
$negativeColor: rgb(230, 6, 6);
.app-main,
.app-main > .row {
  min-height: 100%;
}
.content {
  padding-bottom: 2vh;
  border: #000 1px solid;
  border-color: rgba(0, 0, 0, 0.2);
  border-radius: 3px;
  box-shadow: 0.5px 0.5px 0.5px 0.5px rgba(94, 94, 94, 0.2);
}
.content > .row {
  padding: 1vh 1vw 0;
}
.chart {
  padding: 0 2vw 1vh;
}
.disks {
  background-color: #7600c4;
}
.details {
  background-color: #dfa300;
}
.tabs-item {
  border-radius: 3px;
}
.tabs-item:hover {
  background-color: $main-color;
}
.tabs-item > a {
  color: #000;
}
.tabs-item.active {
  background-color: $main-color;
}
.tabs-item.active > a {
  color: #fff;
}
</style>
