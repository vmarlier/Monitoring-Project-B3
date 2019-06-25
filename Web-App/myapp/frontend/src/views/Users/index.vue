<template>
  <section class="app-main">
    <transition name="fade" mode="out-in">
      <div class="row">
        <div class="col-12">
          <p class="text-left m-0 mt-2">{{ totalUsers }} users</p>
          <div class="row align-items-center">
            <input
              v-model="searchedUser"
              class="col-3 offset-1 form-control mr-sm-2 m-2"
              type="search"
              placeholder="Search a user"
              aria-label="Search"
            />
            <p class="m-0 mr-2 pl-4 font-weight-bold">Only sudoers</p>
            <div class="material-switch text-left">
              <input id="switchSudoers" v-model="isChecked" name="switchSudoers" type="checkbox" />
              <label for="switchSudoers" class="label-switch"></label>
            </div>
          </div>
          <div v-if="filteredList && filteredList.length">
            <paginate :list="filteredList" :per="selectedPerPage" name="users">
              <table class="table table-hover">
                <thead>
                  <tr>
                    <th scope="col">User ID</th>
                    <th scope="col">Username</th>
                    <th scope="col">Display Name</th>
                    <th scope="col">Home directory</th>
                    <th scope="col">Group ID</th>
                    <th scope="col">Sudoer</th>
                    <th scope="col"></th>
                  </tr>
                </thead>
                <tbody>
                  <tr v-for="element in paginated('users')" :key="element.Username">
                    <td>{{ element.UserID }}</td>
                    <td>{{ element.Username }}</td>
                    <td>{{ element.DisplayName }}</td>
                    <td>{{ element.HomeDir }}</td>
                    <td>{{ element.GroupID }}</td>
                    <td>
                      <font-awesome-icon v-if="element.Sudoer" icon="check" class="fa-icon" />
                      <button v-else class="btn btn-sm btn-outline-success" @click.prevent="addSudoer(element.UserID)">
                        Add to sudoers
                      </button>
                    </td>
                    <td>
                      <button
                        v-if="element.Sudoer"
                        class="btn btn-sm btn-outline-danger"
                        @click.prevent="removeSudoer(element.UserID)"
                      >
                        Remove from sudoers
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
              for="users"
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
          <div v-else>No user found</div>
        </div>
      </div>
    </transition>
  </section>
</template>

<script>
import { addSudoer, getUsers, removeSudoer } from '@/api/users'
import PulseLoader from 'vue-spinner/src/PulseLoader.vue'
import config from '@/config'

export default {
  name: 'Users',
  components: { PulseLoader },
  data() {
    return {
      isChecked: false,
      listPerPage: config.PAGINATION_PER_PAGE,
      loading: true,
      paginate: ['users'],
      searchedUser: '',
      selectedPerPage: config.PAGINATION_PER_PAGE[0],
      usersList: []
    }
  },
  computed: {
    filteredList() {
      if (this.isChecked) {
        return this.searchedList.filter(el => el.Sudoer)
      }
      return this.searchedList
    },
    searchedList() {
      const search = this.searchedUser.toLowerCase()
      return this.usersList.filter(
        el =>
          el.Username.toLowerCase().includes(search) ||
          el.DisplayName.toLowerCase().includes(search) ||
          el.UserID.toString().includes(search)
      )
    },
    totalUsers() {
      return this.usersList.length
    }
  },
  created() {
    this.getUsersList()
  },
  methods: {
    addSudoer(username) {
      addSudoer(username).then(response => {
        if (response && response.code && response.code === 200) {
          this.getUsersList()
        }
      })
    },
    getUsersList() {
      getUsers()
        .then(response => {
          if (response && response.data) {
            this.usersList = response.data
          }
        })
        .finally(() => {
          this.loading = false
        })
    },
    removeSudoer(uid) {
      removeSudoer(uid).then(response => {
        if (response && response.code && response.code === 200) {
          this.getUsersList()
        }
      })
    }
  }
}
</script>

<style lang="scss" scoped>
@import '@/style/_variables.scss';

.fa-icon {
  color: $main-color;
}
.label-switch {
  color: $main-color;
  background-color: $main-color;
}
.material-switch > input[type='checkbox'] {
  display: none;
}

.material-switch > label {
  cursor: pointer;
  height: 0px;
  position: relative;
  width: 40px;
}

.material-switch > label::before {
  background: rgb(0, 0, 0);
  box-shadow: inset 0px 0px 10px rgba(0, 0, 0, 0.5);
  border-radius: 8px;
  content: '';
  height: 16px;
  margin-top: -8px;
  position: absolute;
  opacity: 0.3;
  transition: all 0.4s ease-in-out;
  width: 40px;
}
.material-switch > label::after {
  background: rgb(255, 255, 255);
  border-radius: 16px;
  box-shadow: 0px 0px 5px rgba(0, 0, 0, 0.3);
  content: '';
  height: 24px;
  left: -4px;
  margin-top: -8px;
  position: absolute;
  top: -4px;
  transition: all 0.3s ease-in-out;
  width: 24px;
}
.material-switch > input[type='checkbox']:checked + label::before {
  background: inherit;
  opacity: 0.5;
}
.material-switch > input[type='checkbox']:checked + label::after {
  background: inherit;
  left: 20px;
}
</style>
