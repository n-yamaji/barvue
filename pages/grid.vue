<template>
  <v-container 
    fluid
    grid-list-md>
    <barcode-list
      :barcodes="barcodesArray"
      :format="selectedFormat"/>
    <v-select
      v-model="selectedFormat"
      :items="supportedFormats"
      box
      label="バーコードの種類"/>
    <v-textarea
      v-model="barcodes"
      box
      label="バーコード（複数可）"
      auto-grow/>
    <v-btn @click="print()">Print</v-btn>
  </v-container>
</template>

<script>
import VueBarcode from '@chenfengyuan/vue-barcode';
import BarcodeList from '~/components/BarcodeList.vue'

export default {
  components: {
    'barcode': VueBarcode,
    BarcodeList
  },
  data: () => ({
    barcodes: 'い・ろ・は・す 555ML PETx24	4902102091862	1,317\n\
綾鷹 2L PETx6	4902102112208	585\n\
からだ巡茶 からだすこやか茶W 350ML PETx24	4902102108072	2,843',
    selectedFormat: 'CODE128',
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
  }),
  computed: {
    barcodesArray() {
      if (this.barcodes == null) { 
        return [];
      }

      var lines = this.barcodes.split('\n');
      return lines.map((line) => {
        var datas = line.split('\t');
        return {
          label: datas[0],
          code: datas[1],
          retailPrice: datas[2] + '（税込）'
        }
      })
    }
  },
  methods: {
    print() {
      window.print();
    }
  }
}
</script>
