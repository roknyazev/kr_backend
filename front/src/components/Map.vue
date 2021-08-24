<template>
  <div id="uavMap" ref="uavMapRef"></div>
  <div id="uavControls">
    <div class="row">
      <div class="column">
        Создание заказа<br>
        <input type="number" v-model="newOrder.from" placeholder="Откуда"><br>
        <input type="number" v-model="newOrder.to" placeholder="Куда"><br>
        <input type="number" v-model="newOrder.weight" placeholder="Масса, кг"><br>
        <input type="button" @click="sendNewOrder(newOrder)" value="Отправить">
      </div>
      <div class="column">
        Отслеживание заказа<br>
        <input type="number" v-model="orderTrack_input.orderId" placeholder="Откуда"><br>
        <input type="button" @click="callOrderTrack(orderTrack_input)" value="Отправить">
      </div>
      <div class="column">1</div>
    </div>
  </div>
  <div id="devInfo">
    <button @click="getUavs">test</button>
    {{ devInfo }}
  </div>
</template>

<script>
import "leaflet/dist/leaflet.css";
import L from 'leaflet'
import 'leaflet-canvas-markers/leaflet-canvas-markers.js'
import 'leaflet-rotatedmarker/leaflet.rotatedMarker.js';
import 'leaflet.tilelayer.colorfilter/src/leaflet-tilelayer-colorfilter.min.js'
import axios from 'axios'

let myFilter = ['invert:100%','grayscale:42%','bright:80%','contrast:97%','hue:200deg','saturate:150%'];

//let points = [[51.5, -0.09], [60, 45], [50, 3]];
//var icon = L.icon({iconUrl: 'airplane.png', shadowUrl: 'airplane.png', iconSize: [30, 30], shadowSize: [30, 30], iconAnchor: [15, 15], shadowAnchor: [15, 15], popupAnchor: [15, 15]});

export default {
  data() {
    return {
      uavMap: {},

      devInfo: "",

      newOrder: {
        from: "",
        to: "",
        weight: "",
      },
      orderTrack_input: {
        orderId: ""
      },

      UavsFromApi: {
        raw: null,
        small: {},
        medium: {},
        large: {}
      },

      UavsOnMap: {
        small: {},
        medium: {},
        large: {}
      },

      apiServer: {
        url: "http://roknyazev.engineer/"
      },
      apiDataUavs: []
    }
  },

  mounted() {
    this.uavMap = L.map(this.$refs['uavMapRef'], {preferCanvas: true}).setView([56.010569, 92.852572], 7);

    L.tileLayer.colorFilter('https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png', {
      attribution: 'Используется <a href="https://www.openstreetmap.org/copyright">OpenStreetMap</a>',
      filter: myFilter,
    }).addTo(this.uavMap);
  },


  methods: {
    sendNewOrder(_newOrder) {
      alert(JSON.stringify(_newOrder));
    },


    getUavs() {
      axios
          .get(this.apiServer.url + 'state_binary', { timeout: 2000, headers: {} , responseType: "arraybuffer"})
          .then((response) => {
          this.UavsFromApi.raw = response.data
          this.updateUavs();
          })
          .catch(error => {
          console.log(error);
          })
          .finally(() => (this.loading = false));
    },


    updateUavs() {
      let dataView = new DataView(this.UavsFromApi.raw)
      let length = dataView.getInt32(0, true)

      this.UavsOnMap.medium = [];

      let i = 4
      while (i < length + 4) {
        //let type = dataView.getInt32(i, true)
        i += 4
        let lon = dataView.getFloat32(i, true)
        i += 4
        let lat = dataView.getFloat32(i, true)
        i += 4
        let az = dataView.getFloat32(i, true)
        i += 4


        L.canvasMarker(L.latLng(lat, lon), {
          radius: 20,
          img: {
            url: 'airplane.png',    //image link
            size: [30, 30],     //image size ( default [40, 40] )
            rotate: az,         //image base rotate ( default 0 )
            offset: { x: 0, y: 0 }, //image offset ( default { x: 0, y: 0 } )
          },
        }).addTo(this.uavMap);
        //this.UavsOnMap.medium.push(mrk);
      }
      //this.devInfo = this.UavsOnMap.medium;
    }  
  }
}
</script>

<style scoped>
#uavMap {
  height: 75vh;
  width: 100vw;
}
#uavControls {
  height: 15vh;
  width: 100vw;
}
#devInfo {
  height: 10vh;
  width: 100vw;
  background-color: #eee;
}
.column {
  float: left;
  width: 33.33%;
}
.row:after {
  content: "";
  display: table;
  clear: both;
}
</style>
