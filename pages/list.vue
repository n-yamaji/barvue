<template>
  <v-container 
    fluid
    grid-list-md>
    <v-layout column>
      <v-flex class="print-area">
        <v-data-table
          :headers="headers"
          :items="barcodesArray"
          class="elevation-1"
          hide-headers
          hide-actions>
          <template 
            slot="items" 
            slot-scope="props">
            <td class="column-label">{{ props.item.label }}</td>
            <td class="column-price text-xs-right">{{ props.item.retailPrice }}</td>
            <td class="column-barcode">
              <barcode 
                :value="props.item.code" 
                :options="{ 
                  format: format,
                  fontSize: 8,
                  height: 10,
                  width: 1,
                  valid: function() { return true; }
                }" 
                tag="canvas"/>
            </td>
          </template>
        </v-data-table>
      </v-flex>

      <v-flex class="print-ignore">
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
      </v-flex>
    </v-layout>
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
    selectedFormat: 'EAN13',
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
    ],
  headers: [
      { text: '商品名', value: 'label', width: 50 },
      { text: '価格（税込）', value: 'retailPrice', width: 20 },
      { text: '', value: 'code', width: 30 }
    ],
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
          retailPrice: '¥' + datas[2] + '（税込）'
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

<style scope>
@media print {
  .v-navigation-drawer {
    display: none
  }
  .v-toolbar {
    display: none
  }
  .v-footer {
    display: none
  }
  .print-ignore {
    display: none
  }
  .v-content {
    padding: 0 !important;
    margin: 4px
  }
  .container {
    padding: 0
  }
  table { 
    page-break-inside:auto 
  }
  tr { 
    page-break-inside:avoid; 
    page-break-after:auto 
  }
}

table.v-table tbody td, table.v-table tbody th {
    height: 24px;
}
.column-label {
  width: 60%
}
.column-price {
  width: 20%
}
.column-barcode {
  width: 20%
}
</style>
