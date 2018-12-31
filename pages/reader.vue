<template>
  <div>
    <v-text-field :value="barcode"/>
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
      detecteds: [],
      barcode: ''
    }
  },
  methods: {
    logIt (data) {
      let barcode = data.codeResult.code;
      console.log('detected', barcode);
      this.barcode = barcode;
    }
  }
}
</script>
