<template>
  <div class="container login-container">
    <div :class="{ overlay: loading }" class="d-flex justify-content-center">
      <pulse-loader :loading="loading" color="#11a786"></pulse-loader>
    </div>
    <div class="row login-row">
      <div class="offset-lg-3 offset-md-2 col-lg-6 col-md-8 my-auto text-center align-middle login-box">
        <div class="col-lg-12 login-title pt-5">ADMIN PANEL</div>
        <div class="col-lg-12 login-form">
          <div class="col-lg-12 login-form text-left pt-5">
            <form method="POST" @keydown.enter.prevent.self>
              <div :class="{ invalid: errorsFieldsName.includes('username') }" class="form-group">
                <label class="form-control-label">USERNAME</label>
                <input
                  v-model="loginForm.username"
                  type="text"
                  class="form-control form-input"
                  @blur="checkFieldErrors('username', loginForm.username)"
                />
                <p v-if="errorsFieldsName.includes('username')" class="login-form-error">Username cannot be empty</p>
              </div>
              <div :class="{ invalid: errorsFieldsName.includes('password') }" class="form-group">
                <label for="password" class="form-control-label">PASSWORD</label>
                <input
                  v-model="loginForm.password"
                  type="password"
                  name="password"
                  class="form-control form-input"
                  @blur="checkPasswordFormat(loginForm.password)"
                  @keyup.enter="login"
                />
                <p v-if="errorsFieldsName.includes('password')" class="login-form-error">
                  Password cannot be empty and must contain at least 5 characters
                </p>
              </div>
              <div class="col-lg-12 text-center pb-5 pt-3">
                <button type="submit" class="btn btn-outline-primary mb-1" @click.prevent="login">LOGIN</button>
                <div class="col-8 offset-2">
                  <p v-if="errorsFieldsName.length > 0 || alertMessage" class="main-message login-form-error">
                    {{ alertMessage }}
                  </p>
                </div>
              </div>
            </form>
          </div>
        </div>
        <div class="col-lg-3 col-md-2"></div>
      </div>
    </div>
  </div>
</template>
<script>
import PulseLoader from 'vue-spinner/src/PulseLoader.vue'

export default {
  name: 'Login',
  components: { PulseLoader },
  data() {
    return {
      alertMessage: null,
      errorsFieldsName: [],
      loading: false,
      loginForm: {
        username: '',
        password: ''
      },
      showAlertMessage: false
    }
  },
  methods: {
    addFieldToErrors(label, errorsList) {
      const arr = []
      if (!errorsList.find(item => item === label)) {
        arr.push(label)
      }
      return arr
    },
    checkFieldErrors(fieldLabel, fieldValue) {
      // check if field is empty
      if (!fieldValue || !fieldValue.length) {
        this.errorsFieldsName = this.errorsFieldsName.concat(this.addFieldToErrors(fieldLabel, this.errorsFieldsName))
        return true
      }
      this.errorsFieldsName = this.removeFieldFromErrors(fieldLabel, this.errorsFieldsName)

      if (this.errorsFieldsName.length && this.alertMessage) {
        this.alertMessage = null
      }
      return false
    },
    checkIfErrorsIncludeField(field) {
      return this.errorsFieldsName.includes(field)
    },
    checkPasswordFormat(fieldValue) {
      if (this.checkFieldErrors('password', fieldValue)) {
        return
      }
      // check if the password is longer than 5 characters when it is filled
      if (fieldValue.length && fieldValue.length < 5) {
        this.errorsFieldsName = this.errorsFieldsName.concat(this.addFieldToErrors('password', this.errorsFieldsName))
        return
      }
      this.errorsFieldsName = this.removeFieldFromErrors('password', this.errorsFieldsName)
      return
    },
    login() {
      const requiredFields = {
        username: this.loginForm.username,
        password: this.loginForm.password
      }
      this.checkFieldErrors('username', requiredFields.username)
      this.checkPasswordFormat(requiredFields.password)
      if (this.errorsFieldsName.length > 0) {
        this.alertMessage = 'Please, check if all the fields are properly filled'
        return false
      }

      this.loading = true
      this.$store
        .dispatch('Login', { userInfo: this.loginForm, shouldRememberUser: true })
        .then(response => {
          if (response && response.code === 50000) {
            this.alertMessage = 'Wrong credentials, please try again'
          } else {
            this.alertMessage = null
          }
          this.loading = false
          this.$router.push({ path: '/' })
        })
        .catch(() => {
          this.loading = false
          this.$emit('onError')
        })
      // this.signIn(this.loginForm)
    },
    removeFieldFromErrors(label, errorsList) {
      // remove field from errors
      const index = errorsList.indexOf(label)
      if (index !== -1) {
        errorsList.splice(index, 1)
      }
      return errorsList
    }
  }
}
</script>
<style lang="scss">
.login-container,
.login-row {
  height: 100%;
}

.login-box {
  height: 460px;
  background: #1a2226;
  box-shadow: 0 3px 6px rgba(0, 0, 0, 0.16), 0 3px 6px rgba(0, 0, 0, 0.23);
}

.login-title {
  text-align: center;
  font-size: 30px;
  letter-spacing: 2px;
  font-weight: bold;
  color: #11a786;
}

input[type='text'],
input[type='password'] {
  background-color: #1a2226;
  border: none;
  border-bottom: 2px solid #11a786;
  border-top: 0px;
  border-radius: 0px;
  font-weight: bold;
  outline: 0;
  padding-left: 0px;
  color: #ecf0f5;
}

.invalid > input {
  border-bottom: 2px solid #cf0000 !important;
}

.form-group {
  margin-bottom: 40px;
  outline: 0px;
}

.form-control:focus {
  border-color: inherit;
  -webkit-box-shadow: none;
  box-shadow: none;
  border-bottom: 2px solid #11a786;
  outline: 0;
  background-color: #1a2226;
  color: #ecf0f5;
}

input:focus {
  outline: none;
  box-shadow: 0 0 0;
}

label {
  margin-bottom: 0px;
}

.form-control-label {
  font-size: 10px;
  color: #6c6c6c;
  font-weight: bold;
  letter-spacing: 1px;
}

.btn-outline-primary {
  border-color: #11a786;
  color: #11a786;
  border-radius: 0px;
  font-weight: bold;
  letter-spacing: 1px;
  box-shadow: 0 1px 3px rgba(0, 0, 0, 0.12), 0 1px 2px rgba(0, 0, 0, 0.24);
}

.btn-outline-primary:hover {
  background-color: #00cea1;
  border-color: #11a786;
  right: 0px;
}

.login-btm {
  float: left;
}

.login-button {
  padding-right: 0px;
  margin-bottom: 25px;
}

.login-form-error {
  position: absolute;
  font-size: 13px;
  color: #cf0000;
}
.main-message.login-form-error {
  margin-left: auto;
  margin-right: auto;
  left: 0;
  right: 0;
}
</style>
