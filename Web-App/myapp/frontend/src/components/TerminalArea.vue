<template>
  <div class="terminal" draggable="true">
    <textarea
      id="terminal"
      ref="terminal"
      autocorrect="off"
      spellcheck="false"
      class="terminal-input"
      autofocus
      @focus="goToEnd"
      @click.prevent="goToEnd"
      @keydown="handleKey"
    ></textarea>
  </div>
</template>

<script>
import { sendCommand } from '@/api/terminal'

export default {
  name: 'TerminalArea',
  data() {
    return {
      indexHistory: 0,
      history: this.$store.getters.history,
      multiLinesCommand: ''
    }
  },
  mounted() {
    this.$refs.terminal.value =
      "*** An SSH connection is used to access the server. Please use 'ls -ial' instead of 'cd' or 'ls' ***\nadmin $ "
  },
  methods: {
    goToEnd(event) {
      const element = document.getElementById('terminal')
      element.selectionStart = this.$refs.terminal.value.length
    },
    handleKey(event) {
      const text = this.$refs.terminal.value.split('\n')
      const currentLine = this.multiLinesCommand.length ? text[text.length - 1] : text[text.length - 1].slice(8)
      const element = document.getElementById('terminal')
      // scroll to bottom of textarea
      element.scrollTop = element.scrollHeight
      // set cursor at end of textarea
      element.selectionStart = this.$refs.terminal.value.length
      /* Keys to prevent directly */
      const preventedKeyCodes = [13, 38, 40]
      if (preventedKeyCodes.includes(event.keyCode)) {
        event.preventDefault()
      }

      /* Ctrl + a */
      if (event.keyCode === 65 && event.ctrlKey) {
        event.preventDefault()
      }

      /* Ctrl + c */
      if (event.keyCode === 67 && event.ctrlKey) {
        this.$refs.terminal.value += '\nadmin $ '
        return
      }

      /* Enter */
      if (event.keyCode === 13) {
        // reset index history
        this.indexHistory = 0
        // ignore following duplicates
        if (currentLine !== this.history[this.history.length - 1]) {
          // add current line to history
          this.$store.dispatch('addToHistory', currentLine)
        }

        // handle clear command
        if (currentLine === 'clear') {
          this.$refs.terminal.value = 'admin $ '
          return
        }

        // handle multilines commands
        if (currentLine[currentLine.length - 1] === '\\') {
          this.multiLinesCommand += currentLine.slice(0, -1)
          this.$refs.terminal.value += '\n '
          return
        }

        let commandToSend = ''
        // send command
        if (currentLine[currentLine.length - 1] !== '\\') {
          if (this.multiLinesCommand && this.multiLinesCommand.length) {
            this.multiLinesCommand += ` ${currentLine}`
            commandToSend = this.multiLinesCommand
            this.multiLinesCommand = ''
          } else {
            commandToSend = currentLine
          }
          sendCommand(commandToSend)
            .then(response => {
              if (response && response.data && response.data.Body) {
                this.$refs.terminal.value += '\n' + response.data.Body
              }
            })
            .finally(() => {
              this.$refs.terminal.value += '\nadmin $ '
              return
            })
        } else {
          this.$refs.terminal.value += '\nadmin $ '
        }
        return
      }

      /* Backspace */
      if (event.keyCode === 8 && element.selectionStart === this.$refs.terminal.value.length - currentLine.length) {
        event.preventDefault()
        return
      }

      /* Left arrow */
      if (event.keyCode === 37 && element.selectionStart === this.$refs.terminal.value.length - currentLine.length) {
        event.preventDefault()
        return
      }

      /* Up arrow */
      if (
        event.keyCode === 38 &&
        (currentLine.length || element.selectionStart === this.$refs.terminal.value.length) &&
        this.indexHistory < this.history.length
      ) {
        if (currentLine.length) {
          this.$refs.terminal.value = this.$refs.terminal.value.slice(0, -currentLine.length)
        }
        this.$refs.terminal.value += this.history[this.history.length - this.indexHistory - 1]
        this.indexHistory++
        return
      }

      /* Down arrow */
      if (
        event.keyCode === 40 &&
        (currentLine.length || element.selectionStart === this.$refs.terminal.value.length) &&
        this.indexHistory <= this.history.length &&
        this.indexHistory > 1
      ) {
        if (currentLine.length) {
          this.$refs.terminal.value = this.$refs.terminal.value.slice(0, -currentLine.length)
        }
        this.indexHistory--
        this.$refs.terminal.value += this.history[this.history.length - this.indexHistory]
        return
      }
    }
  }
}
</script>

<style lang="scss" scoped>
.app-main {
  height: 100%;
}

.terminal {
  position: relative;
  background-color: #111;
  height: 100%;
  width: 100%;

  -webkit-box-shadow: 0px 13px 49px -24px rgba(0, 0, 0, 0.75);
  -moz-box-shadow: 0px 13px 49px -24px rgba(0, 0, 0, 0.75);
  box-shadow: 0px 13px 49px -24px rgba(0, 0, 0, 0.75);
}

.terminal-input {
  height: 100%;
  width: 100%;
  background-color: #111;
  padding: 10px 0 10px 10px;
  border: 0;
  color: #fff;
  font-size: 14px;
  -webkit-box-sizing: border-box;
  -moz-box-sizing: border-box;
  box-sizing: border-box;
}

.terminal-input,
.terminal-input:focus {
  outline: none;
  resize: none;
  border-radius: 5px;
  font-family: Courier;
}

textarea::-webkit-input-placeholder {
  width: 10px;
  text-shadow: none;
  -webkit-text-fill-color: #f00;
}

.brand,
.brand:visited {
  position: absolute;
  padding: 0;
  margin: 0;
  bottom: 5px;
  right: 10px;
  text-decoration: none;
  color: rgba(0, 0, 0, 0.4);
  font-size: 14px;
  font-family: 'Pacifico', cursive;
}

.brand:hover {
  text-decoration: underline;
}

/* --- Codecademy preview ---- */
@media screen and (max-width: 300px) and (max-height: 200px) {
  .brand,
  .brand:visited {
    display: none;
  }

  .terminal {
    padding: 0;
    height: 80px;
    width: 230px;
  }
}
</style>
