<template>
  <v-container>
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

        <v-card height="480px">
          <div class="card-title">
            <v-card-title primary-title>
              <div>
                <div class="item-name">{{ barcode.label }}</div>
              </div>
            </v-card-title>
          </div>

          <div class="card-footer">
            <div class="card-atrributes">
              <v-card-text>
                <div>
                  <span class="price">¥{{ barcode.retailPrice }}</span>
                </div>
              </v-card-text>
            </div>

            <div class="text-xs-center card-barcode">
              <barcode 
                :value="barcode.code" 
                :options="{ 
                  format: format,
                  width: '3px',
                  valid: function() { return true; }
                }" 
                tag="canvas"/>
            </div>
          </div>
        </v-card>
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
      default: 'CODE128'
    },
    barcodes: {
      type: Object,
      default: null
    }
  }
}
</script>

<style>
.card-footer {
  width: 100%;
  position: absolute;
  bottom: 0;
  margin:0 auto;
  margin-right: 0px;
  margin-bottom: 12px;
}
.card-atrributes {
  text-align: right;
}
.item-name {
  font-size: 40px;
}
.price {
  font-size: 40px;
}
</style>
