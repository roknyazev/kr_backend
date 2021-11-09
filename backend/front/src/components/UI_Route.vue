<template>
  <div v-if="wrong_id !== undefined">
    <v-alert
      v-if="!wrong_id"
      border="left"
      colored-border
      type="warning"
      elevation="2"
      icon="mdi-alert-circle-outline"
    >
      <div class="text-body-1">Заказа с таким номером не существует</div>
    </v-alert>
    <div v-else-if="wrong_id">

    <v-card>
      <v-card-title class="subheading font-weight-bold"
        >Заказ #{{ order_id }}</v-card-title
      >

      <v-divider></v-divider>

      <v-list dense>
        <v-list-item>
          <v-list-item-content class="text-body-1"
            >Масса груза</v-list-item-content
          >
          <v-list-item-content class="align-end text-body-1"
            >{{ weight }} кг</v-list-item-content
          >
        </v-list-item>
      </v-list>
    </v-card>

    <v-timeline id="uavRoute" align-top dense>
      <div v-for="point in route" :key="point">
        <!--<div v-if="point === route[0]">first {{point}}</div>
        <div v-else-if="point === route[route.length -1]">last {{point}}</div>-->
        <v-timeline-item v-if="point === route[0]" color="primary" dense>
          <template v-slot:icon>
            <v-icon dense color="white">mdi-airplane-takeoff</v-icon>
          </template>
          <div class="text-h6">
            <nobr>Хаб {{ point.HubID }}</nobr>
          </div>
          <nobr class="text-subtitle-2"
            ><v-icon dense :dark="appDark">mdi-map-marker-outline</v-icon>
            <a
              @click="
                focusAtHub(
                  (point.lan_pos * 180) / Uav_PI,
                  (point.lon_pos * 180) / Uav_PI
                )
              "
            >
              {{ ((point.lan_pos * 180) / Uav_PI).toFixed(6) }},
              {{ ((point.lon_pos * 180) / Uav_PI).toFixed(6) }}</a
            ></nobr
          >
        </v-timeline-item>
        <v-timeline-item
          v-else-if="point === route[route.length - 1]"
          color="white"
          dense
        >
          <template v-slot:icon>
            <v-icon dense color="black">mdi-airplane-landing</v-icon>
          </template>
          <div class="text-h6">
            <nobr>Хаб #{{ point.HubID }}</nobr>
          </div>
          <nobr class="text-subtitle-2"
            ><v-icon dense dark>mdi-map-marker-outline</v-icon>
            <a
              @click="
                focusAtHub(
                  (point.lan_pos * 180) / Uav_PI,
                  (point.lon_pos * 180) / Uav_PI
                )
              "
              style="[{ 'white': appDark }, 'black']"
            >
              {{ ((point.lan_pos * 180) / Uav_PI).toFixed(6) }},
              {{ ((point.lon_pos * 180) / Uav_PI).toFixed(6) }}</a
            ></nobr
          >
        </v-timeline-item>
        <v-timeline-item v-else color="grey darken-2" dense small>
          <div class="text-bold">
            <nobr>Хаб {{ point.HubID }}</nobr>
          </div>
          <nobr class="text-subtitle-2">
            <v-icon dense :dark="appDark">mdi-map-marker-outline</v-icon>
            <a
              @click="
                focusAtHub(
                  (point.lan_pos * 180) / Uav_PI,
                  (point.lon_pos * 180) / Uav_PI
                )
              "
              style="[{ 'white': appDark }, 'black']"
            >
              {{ ((point.lan_pos * 180) / Uav_PI).toFixed(6) }},
              {{ ((point.lon_pos * 180) / Uav_PI).toFixed(6) }}</a
            ></nobr
          >
        </v-timeline-item>
      </div>
    </v-timeline>
    </div>
  </div>
</template>
<script>
import L from "leaflet";

export default {
  props: {
    wrong_id: Boolean,
    appDark: Boolean,
    departure_lat: Number,
    departure_lon: Number,
    arrival_lat: Number,
    arrival_lon: Number,
    weight: String,
    order_id: String,
    route: [],
    map: {},
  },
  data() {
    return {
      Uav_PI: 3.141592,
      Hub_CircleMarker: Object,
    };
  },
  methods: {
    focusAtHub(lat, lon) {
      this.map.flyTo(new L.LatLng(lat, lon), 10);
    },
  },
};
//this.map.flyTo((point.lan_pos * 180) / Uav_PI, (point.lon_pos * 180) / Uav_PI, 8)
</script>
<style>
a:hover {
  text-decoration: underline;
}
</style>
