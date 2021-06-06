<template>
  <v-container class="fill-height" width="800px">
    <v-row
      align="center"
      justify="center"
    >
      <v-col class="text-center">
        <v-form
            ref="form"
            v-model="valid"
            lazy-validation
        >
          <v-text-field
            v-model="userid"
            :counter="10"
            :rules="requiredRule"
            label="ID"
            required
          ></v-text-field>

          <v-text-field
            v-model="password"
            :rules="requiredRule"
            type="password"
            label="Password"
            required
          ></v-text-field>

          <v-btn
            color="primary"
            large
            @click="login"
          >
          はじめる
          </v-btn>
        </v-form>
      </v-col>
    </v-row>
    <v-row><v-col class="text-center"><a href="https://scrimba.com/playlist/pP4xZu3">lets study vuetify!</a></v-col></v-row>
  </v-container>
</template>
<script>
import axios from 'axios'
export default {
  name: 'TitlePage',
  data: function () {
    return {
      userid: '',
      password: '',
      requiredRule: [
        v => !!v || 'required property'
      ],
      valid: false
    }
  },
  methods: {
    async login () {
      try {
        const instance = axios.create({})
        await instance.post('/v1/auth/login/', {
          userID: this.userid,
          password: this.password
        })
        this.$router.push('/images')
      } catch (error) {
        if (error.response.status === 401) {
          window.alert('IDもしくはパスワードが不正です')
        } else {
          console.log(error)
        }
      }
    }
  }
}
</script>
