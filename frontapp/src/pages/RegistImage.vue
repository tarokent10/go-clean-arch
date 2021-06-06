<template>
  <v-form ref="form">
    <v-container class="fill-height content">
      <v-card width=800 class="mx-auto mt-6">
        <v-card-text>
          <v-row
            align="center"
            justify="center"
          >
            <v-col class="text-center">
              <v-file-input
                show-size
                placeholder="Upload your favorite picture"
                label="my picture"
                @change="fileinput"
                :rules="required"
              >
                <template v-slot:selection="{ text }">
                  <v-chip
                    small
                    label
                    color="primary"
                  >
                    {{ text }}
                  </v-chip>
                </template>
              </v-file-input>
            </v-col>
          </v-row>
          <v-row class="mx-5">
            <v-col>
              <v-text-field
                v-model="name"
                label="name"
                clearable
                :rules="required">
              </v-text-field>
            </v-col>
          </v-row>
          <v-row class="mx-5">
            <v-col>
              <v-textarea
                v-model="description"
                label="description"
                :rules="required">
              </v-textarea>
            </v-col>
          </v-row>
          <v-row>
            <v-col>
              <v-img
                v-bind:src="imgUri"
              ></v-img>
            </v-col>
          </v-row>
          <v-row class="mx-5">
            <v-col>
              <v-btn block color="primary" @click="regist" :loading="loading">登録する
              </v-btn>
            </v-col>
          </v-row>
        </v-card-text>
      </v-card>
    </v-container>
  </v-form>
</template>
<script>
import axios from 'axios'
export default {
  name: 'RegistImage',
  data: function () {
    return {
      name: '',
      description: '',
      imgUri: '',
      loading: false,
      required: [
        v => !!v || 'required property'
      ]
    }
  },
  methods: {
    fileinput: function (f) {
      const changeImgfunc = this.changeImg
      const reader = new FileReader()
      reader.onload = function (e) {
        if (e != null) {
          changeImgfunc(e.target.result)
        }
      }
      reader.readAsDataURL(f)
    },
    changeImg: function (uri) {
      this.imgUri = uri
    },
    regist: function () {
      const valid = this.$refs.form.validate()
      if (!valid) {
        return
      }
      this.loading = true
      // const as = axios.create({
      //   baseURL: process.env.VUE_APP_API_HOST,
      //   headers: {
      //     'Content-Type': 'application/json'
      //   },
      //   responseType: 'json'
      // })
      axios.post('/v1/images/', {
        name: this.name,
        data: this.imgUri,
        description: this.description
      }, {
        withCredentials: true
      }).then(function (response) {
        console.log('post sucessed')
      })
        .catch(function (error) {
          console.log(error)
        })
      this.loading = false
    }
  }
}
</script>
<style scoped>
.content {
  max-width: 1000px
}
</style>
