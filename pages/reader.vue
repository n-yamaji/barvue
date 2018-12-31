<template>
  <div>
    <quagga-scanner 
      :onDetected="logIt" 
      :readerSize="readerSize" 
      :readerType="'ean_reader'" />
  </div>
</template>

<script>
import { QuaggaScanner } from 'vue-quaggajs'

export default {

  components: {
    QuaggaScanner
  },
  data () {
    return {
      readerSize: {
        width: 640,
        height: 480
      },
      detecteds: []
    }
  },
  methods: {
    logIt (data) {
      console.log('detected', data)
    }
  }
}
</script>
