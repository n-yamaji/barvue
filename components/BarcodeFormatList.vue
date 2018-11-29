<template>
  <v-container>
    <v-layout column>
      <v-flex>
        <v-text-field
          v-model="barcodeValue"
          label="barcode"
          box/>
      </v-flex>

      <v-flex>
        <v-layout 
          row 
          wrap
          grid-list-xl>
          <v-flex 
            v-for="f in formats" 
            :key="f"
            xs6
            sm6
            md6>
            <p>{{ f }}</p>
            <barcode 
              :value="barcodeValue" 
              :options="{ 
                format: f,
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
    formats: {
      type: Array,
      required: true
    }
  },
  data: () => ({
    barcodeValue: '12345'
  })
}
</script>
