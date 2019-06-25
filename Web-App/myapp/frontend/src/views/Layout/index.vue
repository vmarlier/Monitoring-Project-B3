<template>
  <div class="main">
    <nav class="navbar navbar-expand-sm bg-dark navbar-dark fixed-top">
      <div class="navbar-title">
        <h3>{{ $route.name }}</h3>
      </div>

      <ul class="navbar-nav ml-auto">
        <div class="btn-group">
          <a class="btn p-3" data-toggle="dropdown" href="#">
            <font-awesome-icon icon="user" class="fa-icon" />
          </a>
          <ul class="dropdown-menu dropdown-menu-right">
            <li class="nav-item">
              <a class="dropdown-item" href="#" @click.prevent="logout">Logout</a>
            </li>
          </ul>
        </div>
      </ul>
    </nav>
    <div class="row p-0 m-0 h-100">
      <nav id="sidebar-container" class="content-left-sidebar col-md-2 d-none d-md-block bg-dark sidebar">
        <ul class="nav flex-column list-group sticky-top sticky-offset">
          <navigation-link name="Dashboard" route="/" />
          <navigation-link name="Processes" route="/processes" />
          <navigation-link name="Users" route="/users" />
          <navigation-link name="Terminal" route="/terminal" />
        </ul>
      </nav>
      <div :class="{ 'terminal-bg': isFullWidth }" class="col-md-10 app content app-content container-fluid">
        <div class="content-wrapper">
          <router-view />
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import NavigationLink from '@/components/NavigationLink'

export default {
  name: 'Layout',
  components: { NavigationLink },
  computed: {
    isFullWidth() {
      return this.$route.name === 'Terminal'
    }
  },
  methods: {
    logout() {
      this.$store.dispatch('LogOut').then(() => {
        location.reload()
      })
    }
  }
}
</script>

<style lang="scss">
@import '@/style/_variables.scss';
.navbar {
  border-bottom: $main-color 2px solid;
}
body {
  padding-top: 76px;
}
.content-wrapper {
  height: 100%;
}
.navbar-title {
  width: 100%;
  color: #fff;
}
.navbar .dropdown-menu {
  margin: 0;
}
.terminal-bg {
  background-color: #111;
  padding: 0;
  padding-bottom: 0;
  border-radius: 0;
}
.sidebar .nav-item:hover,
.navbar .dropdown-item:hover {
  background-color: $main-color;
  border-color: $main-color;
  color: #fff;
}

.main {
  height: 100%;
}
#sidebar-container {
  padding: 0;
  position: sticky;
}
a.router-link-exact-active {
  background-color: $main-color;
}
</style>
