<template>
  <v-container class="fill-height" fluid>
    <v-row justify-content="start">
      <v-col v-for="(image, idx) in images" :key=idx>
        <v-card width="200px" height="auto" elevation="5" >
          <v-img
            v-bind:src="image.uri"
            class="white--text align-end"
            gradient="to bottom, rgba(0,0,0,.1), rgba(0,0,0,.5)"
          >
          </v-img>
          <p>氏名：{{ image.name }} </p>
          <p>部署：{{ image.department }} </p>
          <p>備考：{{ image.remarks }}</p>
        </v-card>
      </v-col>
    </v-row>
  </v-container>
</template>
<script>
import axios from 'axios'
export default {
  name: 'images',
  data: function () {
    return {
      images: []
    }
  },
  methods: {
    get: function () {
      const self = this
      axios.get('/v1/employees/').then(function (res) {
        res.data.forEach(e => {
          const d = 'data:image/png;base64,' + e.picture
          self.images.push({ // 力技
            name: e.name,
            uri: d,
            updateDateTime: e.updateDateTime,
            department: e.department,
            remarks: e.remarks
          })
        })
      })
        .catch(function (error) {
          console.log(error)
        })
    }
  },
  created: function () {
    this.get()
  }
}
</script>
<style scoped>
</style>
