<template>
  <v-container>
    <v-layout column>
      <v-flex>
        <v-layout 
          row 
          wrap
          justify-space-between
          grid-list-xl>
          <v-flex 
            v-for="barcode in barcodes" 
            :key="barcode.code"
            xs6
            sm6
            md6>
            <p>{{ barcode.label }}</p>
            <p>¥{{ barcode.retailPrice }}</p>
            <barcode 
              :value="barcode.code" 
              :options="{ 
                format: format,
                valid: function() { return true; }
              }" 
              tag="canvas"/>
          </v-flex>
        </v-layout>
      </v-flex>
    </v-layout>
  </v-container>
</template>

<script>
import VueBarcode from '@chenfengyuan/vue-barcode';

export default {
  components: {
    'barcode': VueBarcode
  },
  props: {
    format: {
      type: String,
      default: 'EAN13'
    },
    barcodes: {
      type: Object,
      default: null
    }
  }
}
</script>
