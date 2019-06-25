<template>
  <section class="app-main">
    <transition name="fade" mode="out-in">
      <div class="row">
        <div class="col-12">
          <p class="text-left m-0 mt-2">{{ totalProcesses }} processes</p>
          <div class="row align-items-center">
            <input
              v-model="searchedProcess"
              class="col-3 offset-1 form-control mr-sm-2 m-2"
              type="search"
              placeholder="Search a process"
              aria-label="Search"
            />
          </div>
          <div v-if="filteredList && filteredList.length">
            <paginate :list="filteredList" :per="selectedPerPage" name="processes">
              <table class="table table-hover">
                <thead>
                  <tr>
                    <th scope="col">PID</th>
                    <th scope="col">PPID</th>
                    <th scope="col">Name</th>
                    <th scope="col"></th>
                  </tr>
                </thead>
                <tbody>
                  <tr v-for="element in paginated('processes')" :key="element.Pid">
                    <td>{{ element.Pid }}</td>
                    <td :title="getPPid(element.PPid)">{{ element.PPid }}</td>
                    <td>{{ element.Executable }}</td>
                    <td>
                      <button class="btn btn-sm btn-outline-danger" @click.prevent="killProcess(element.Pid)">
                        Kill
                      </button>
                    </td>
                  </tr>
                </tbody>
              </table>
            </paginate>
            <paginate-links
              :hide-single-page="true"
              :show-step-links="true"
              :limit="3"
              :step-links="{
                next: 'Next',
                prev: 'Previous'
              }"
              for="processes"
            ></paginate-links>
            <div v-if="filteredList && filteredList.length > perPage" class="btn-group">
              <button
                type="button"
                class="btn btn-outline-dark dropdown-toggle btn-sm"
                data-toggle="dropdown"
                aria-haspopup="true"
                aria-expanded="false"
              >
                {{ selectedPerPage }}
              </button>
              <div class="dropdown-menu">
                <a
                  v-for="perPage in listPerPage"
                  :key="perPage"
                  :class="{ active: perPage === selectedPerPage }"
                  class="dropdown-item"
                  @click.prevent="selectedPerPage = perPage"
                >
                  {{ perPage }}
                </a>
              </div>
            </div>
          </div>
          <div v-else>No process found</div>
        </div>
      </div>
    </transition>
  </section>
</template>

<script>
import { getProcesses, killProcess } from '@/api/processes'
import PulseLoader from 'vue-spinner/src/PulseLoader.vue'
import config from '@/config'

export default {
  name: 'Processes',
  components: { PulseLoader },
  data() {
    return {
      listPerPage: config.PAGINATION_PER_PAGE,
      loading: false,
      paginate: ['processes'],
      processesList: [],
      searchedProcess: '',
      selectedPerPage: config.PAGINATION_PER_PAGE[0]
    }
  },
  computed: {
    filteredList() {
      const search = this.searchedProcess.toLowerCase()
      return this.processesList.filter(
        el => el.Executable.toLowerCase().includes(search) || el.Pid.toString().includes(search)
      )
    },
    totalProcesses() {
      return this.processesList.length
    }
  },
  created() {
    this.getProcessesList()
  },
  mounted() {
    this.$nextTick(function() {
      this.interval = window.setInterval(() => {
        this.getProcessesList()
      }, 1000)
    })
  },
  beforeDestroy() {
    clearInterval(this.interval)
  },
  methods: {
    getPPid(ppid) {
      const elementFound = this.processesList.find(el => el.Pid === ppid)
      if (elementFound) {
        return `${elementFound.Executable}`
      }
    },
    getProcessesList() {
      this.loading = true
      getProcesses()
        .then(response => {
          if (response && response.data && response.data.Procs) {
            this.processesList = response.data.Procs
          }
        })
        .finally(() => {
          this.loading = false
        })
    },
    killProcess(pid) {
      killProcess(pid).then(response => {
        if (response && response.code && response.code === 200) {
          this.getProcessesList()
        }
      })
    }
  }
}
</script>

<style lang="scss" scoped>
.fa-icon {
  color: #000;
}
</style>
