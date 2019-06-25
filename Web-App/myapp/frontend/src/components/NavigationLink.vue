<template>
  <li class="nav-item">
    <router-link :to="route" class="nav-link p-3">
      <span v-if="name && name.length">{{ name }}</span>
      <font-awesome-icon v-if="icon && icon.length" :icon="icon" class="fa-icon" />
    </router-link>
  </li>
</template>

<script>
export default {
  name: 'NavigationLink',
  props: {
    icon: { type: String, default: '' },
    name: { type: String, default: '' },
    route: { type: String, required: true }
  }
}
</script>

<style lang="scss">
.nav-link,
.fa-icon {
  color: #fff;
}
.nav-link:hover .fa-icon {
  color: #bababa;
}

.nav-link:hover,
a:hover {
  text-decoration: none;
  color: #fff;
}
</style>
