<template>
  <section class="app-main">
    <transition name="fade" mode="out-in">
      <terminal-area></terminal-area>
    </transition>
  </section>
</template>

<script>
import TerminalArea from '@/components/TerminalArea'
export default {
  name: 'Terminal',
  components: { TerminalArea }
}
</script>

<style lang="scss" scoped>
.app-main {
  height: 100%;
}
</style>
