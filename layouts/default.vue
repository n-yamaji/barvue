<template>
  <v-app light>
    <v-navigation-drawer
      mini-variant
      permanent
      fixed
      app>
      <v-list>
        <v-list-tile
          v-for="(item, i) in items"
          :to="item.to"
          :key="i"
          router
          exact>
          <v-list-tile-action>
            <v-icon v-html="item.icon" />
          </v-list-tile-action>
          <v-list-tile-content>
            <v-list-tile-title v-text="item.title" />
          </v-list-tile-content>
        </v-list-tile>
      </v-list>
    </v-navigation-drawer>

    <v-toolbar
      fixed
      app>
      <v-toolbar-title v-text="title"/>
    </v-toolbar>
    <v-content>
      <v-container>
        <nuxt />
      </v-container>
    </v-content>
    <v-footer
      fixed
      app
    >
      <span>&copy; 2018</span>
    </v-footer>
  </v-app>
</template>

<script>
  export default {
    data() {
      return {
        items: [
          { icon: 'grain', title: 'Pattern', to: '/' },
          { icon: 'grid_on', title: 'Grid', to: '/grid' },
          { icon: 'list', title: 'List', to: '/list' }
        ],
        title: 'barvue'
      }
    }
  }
</script>
