<template>
  <barcode-format-list :formats="supportedFormats"/>
</template>

<script>
import VueBarcode from '@chenfengyuan/vue-barcode';
import BarcodeFormatList from '~/components/BarcodeFormatList.vue'

export default {

  components: {
    'barcode': VueBarcode,
    BarcodeFormatList
  },
  data: () => ({
    supportedFormats: [
      'CODE128', 
      'CODE128A', 
      'CODE128B', 
      'CODE128C', 
      'EAN13', 
      'EAN8', 
      'EAN5', 
      'EAN2', 
      'UPC',
      'CODE39', 
      'ITF14', 
      'MSI', 
      'MSI10', 
      'MSI11', 
      'MSI1010', 
      'MSI1110', 
      'pharmacode', 
      'codabar', 
    ]
  })
}
</script>
