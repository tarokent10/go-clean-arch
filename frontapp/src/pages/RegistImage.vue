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
                label="アイコン"
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
                label="名前"
                clearable
                :rules="required">
              </v-text-field>
            </v-col>
          </v-row>
          <v-row class="mx-5">
            <v-col>
              <v-text-field
                v-model="department"
                label="部署"
                clearable>
              </v-text-field>
            </v-col>
          </v-row>
          <v-row class="mx-5">
            <v-col>
              <v-textarea
                v-model="remarks"
                label="備考"
                clearable>
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
              <v-btn block color="primary" @click="regist">登録する
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
      remarks: '',
      department: '',
      imgUri: '',
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
      axios.post('/v1/employee/', {
        name: this.name,
        picture: this.imgUri,
        updateDateTime: '',
        department: this.department,
        remarks: this.remarks
      }).then(function (response) {
        window.alert('従業員情報を登録しました')
      })
        .catch(function (error) {
          console.log(error)
        })
    }
  }
}
</script>
<style scoped>
.content {
  max-width: 1000px
}
</style>
