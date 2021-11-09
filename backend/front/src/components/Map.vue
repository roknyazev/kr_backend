<template>
  <div id="app">
    <loader
      v-if="!appLoaded"
      style="z-index: 999"
      object="#ffffff"
      color1="#ffffff"
      color2="#17fd3d"
      size="5"
      speed="1"
      bg="#343a40"
      objectbg="#999793"
      opacity="80"
      disableScrolling="false"
      name="dots"
    ></loader>
    <v-app>
      <v-system-bar window :dark="appDark">
        <v-tooltip bottom style="z-index: 999" open-delay="250">
          <template v-slot:activator="{ on, attrs }">
            <v-switch
              v-model="appDark"
              dense
              small
              @change="changeMapTheme()"
            ></v-switch>
            <v-icon v-bind="attrs" v-on="on">{{
              appDark ? "mdi-lightbulb" : "mdi-lightbulb-outline"
            }}</v-icon>
          </template>
          <span>{{
            appDark ? "Включить светлый интерфейс" : "Включить тёмный интерфейс"
          }}</span>
        </v-tooltip>

        <nobr style="overflow: hidden">{{ infoMessage }}</nobr>
        <v-spacer></v-spacer>
        <v-tooltip bottom style="z-index: 999" open-delay="500">
          <template v-slot:activator="{ on, attrs }">
            <nobr style="overflow: hidden; min-width: fit-content">
              <span
                class="system-bar-element"
                style="float: right"
                v-bind="attrs"
                v-on="on"
              >
                <v-icon>mdi-lan-connect</v-icon>
                <span>{{ statusBar.connectionSpeed }}</span>
                мс
              </span>
            </nobr>
          </template>
          <span>Время ответа от сервера. Меньше — лучше</span>
        </v-tooltip>
        <v-tooltip bottom style="z-index: 999" open-delay="500">
          <template v-slot:activator="{ on, attrs }">
            <nobr style="overflow: hidden; min-width: fit-content">
              <span
                class="system-bar-element"
                style="float: right"
                v-bind="attrs"
                v-on="on"
              >
                <v-icon>mdi-airplane</v-icon>
                <span
                  >{{ statusBar.totalUavs }}
                  ЛА
                </span>
              </span>
            </nobr>
          </template>
          <span>ЛA полете в данный момент</span>
        </v-tooltip>
        <v-tooltip bottom style="z-index: 999" open-delay="500">
          <template v-slot:activator="{ on, attrs }">
            <nobr style="overflow: hidden; min-width: fit-content">
              <span
                class="system-bar-element"
                style="float: right"
                v-bind="attrs"
                v-on="on"
              >
                <v-icon>mdi-bag-checked</v-icon>
                <span
                  >{{ statusBar.totalOrders }}
                  <span v-if="!$isMobile()">активных заказов</span>
                </span>
              </span>
            </nobr>
          </template>
          <span>Активных заказов в данный момент</span>
        </v-tooltip>
      </v-system-bar>
      <v-row no-gutters v-if="!$isMobile()">
        <v-col cols="3" id="uavControlsDesktop">
          <v-card id="uavControlsDesktopCard" outlined tile :dark="appDark">
            <div>
              <v-container>
                <v-card-actions class="justify-center">
                  <v-btn-toggle
                    color="primary"
                    dense
                    :outlined="appDark"
                    mandatory
                    rounded
                    v-model="window"
                    style="display: block; margin: auto"
                    @change="toggleDeliveryZone(); resetAddOrder()"
                  >
                    <v-btn> Добавить </v-btn>
                    <v-btn> Отследить </v-btn>
                  </v-btn-toggle>
                </v-card-actions>

                <v-window v-model="window" class="elevation-1">
                  <v-window-item>
                    <v-card flat>
                      <v-card-title>Создание заказа</v-card-title>
                      <v-card-text>
                        <v-alert
                          v-if="UI.addOrder.error"
                          border="left"
                          colored-border
                          type="error"
                          elevation="2"
                          icon="mdi-alert-circle-outline"
                          prominent
                        >
                          <div class="text-h6">Ошибка при создании заказа</div>
                          <div class="text-body-2">Проверьте нахождение точек в зонах доставки</div>
                        </v-alert>
                        <v-form ref="form" v-model="valid" lazy-validation>
                          <transition name="fade">
                            <template>
                              <v-btn
                                @click="addOrderMarkers"
                                v-if="!UI.addOrder.arePointsAdded"
                                rounded
                                outlined
                                block
                              >
                                <v-icon dense v-if="!UI.addOrder.processing"
                                  >mdi-plus</v-icon
                                >
                                <v-icon dense v-else
                                  >mdi-map-marker-outline</v-icon
                                >
                                &nbsp;{{ UI.addOrder.locBtnText }}
                              </v-btn>
                            </template>
                          </transition>
                          <transition name="fade">
                            <div
                              class="pt-2 text-caption"
                              v-if="UI.addOrder.processing"
                            >
                              на карте, используя правую кнопку мыши (ПКМ)
                            </div>
                          </transition>
                          <v-text-field
                            v-model="UI.addOrder.orderWeight"
                            ref="addOrder_Weight"
                            hint="Укажите массу груза"
                            :counter="10"
                            class="mt-5"
                            color="[{ 'white': appDark }, 'black']"
                            outlined
                            dense
                            label="Масса"
                            suffix="кг"
                            required
                          ></v-text-field>
                          <div class="text-right">
                            <v-btn
                              :disabled="!valid"
                              color="[{ 'white': appDark }, 'black']"
                              :loading="UI.addOrder.sending"
                              rounded
                              outlined
                              class="mt-4"
                              @click="sendNewOrder()"
                            >
                              Создать
                            </v-btn>
                          </div>
                        </v-form>
                      </v-card-text>
                    </v-card>
                  </v-window-item>
                  <v-window-item>
                    <transition name="fade">
                      <v-card flat>
                        <v-card-title>Отслеживание заказа</v-card-title>
                        <v-card-text>
                          <v-form ref="form" lazy-validation>
                            <v-text-field
                              :counter="10"
                              outlined
                              dense
                              color="[{ 'white': appDark }, 'black']"
                              label="Номер заказа"
                              required
                              v-model="UI.trackOrder.input_id"
                              append-icon="mdi-magnify"
                              @click:append="
                                sendTrackOrder(UI.trackOrder.input_id)
                              "
                            ></v-text-field>
                            <!--<v-btn :disabled="!valid" color="success" class="mr-4" @click="UI.watchOrder.success = !UI.watchOrder.success">
                        Отследить
                      </v-btn>-->
                          </v-form>
                        </v-card-text>
                        <v-card-text>
                          <!--{{UI.addOrder.orderInfo}}-->
                          <UI_Route
                            :wrong_id="UI.track.result"
                            :appDark="appDark"
                            :map="map"
                            :order_id="UI.addOrder.orderInfo.order_id"
                            :weight="UI.addOrder.orderInfo.weight"
                            :arrival_lat="UI.addOrder.orderInfo.first_lat"
                            :arrival_lon="UI.addOrder.orderInfo.first_lon"
                            :route="UI.addOrder.orderInfo.route"
                          ></UI_Route>
                        </v-card-text>
                      </v-card>
                    </transition>
                  </v-window-item>
                </v-window>
              </v-container>
            </div>
          </v-card>
          <v-card
            id="uavInfoDesktopCard"
            v-if="UI.uavInfo.show"
            outlined
            tile
            :dark="appDark"
          >
            <v-container>
              <v-card>
                <v-card-title class="mt-8">
                  <v-avatar size="56">
                    <img :src="UI.uavInfo.image" />
                  </v-avatar>
                  <div class="ml-3">
                    ЛА #{{ UI.uavInfo.id }}
                    <div class="text-body-2">{{ UI.uavInfo.type }}</div>
                  </div>
                </v-card-title>
                <v-simple-table dense>
                  <template v-slot:default>
                    <thead>
                      <tr>
                        <th class="text-left">Заказы на борту</th>
                      </tr>
                    </thead>
                    <tbody>
                      <tr
                        v-for="order in UI.uavInfo.ordersOnboard"
                        :key="order"
                      >
                        <td><span class="text-body-1">{{ order }}</span></td>
                        <td>
                        <v-btn
                          class="my-2"
                          fab
                          x-small
                          color="primary"
                          @click="sendTrackOrder(order)"
                          style="display: block; margin-right: 0; margin-left: auto"
                        >
                          <v-icon dark>
                            mdi-map-marker-question
                          </v-icon>
                        </v-btn>
                        </td>
                      </tr>
                      <!--
                      {{
                        UI.uavInfo.ordersOnboard
                      }}-->
                    </tbody>
                  </template>
                </v-simple-table>
              </v-card>
            </v-container>
          </v-card>
        </v-col>

        <v-col cols="9">
          <v-card outlined tile :dark="appDark" style="border: none">
            <div class="uavKartaContainerDesktop">
              <div
                id="uavKarta"
                ref="uavKartaRef"
                :style="[
                  return_is_UI_AddOrder_Proc() ? { cursor: 'crosshair' } : '',
                ]"
              >
                <!--
                <div id="uavTimeline">
                  <v-row>
                    <v-col cols="1"> </v-col>
                    <v-col cols="10">
                      <v-slider
                        v-if="appDark"
                        hint="Таймлайн"
                        max="0"
                        min="-50"
                        color="white"
                        thumb-color="red"
                        track-color="grey darken-3"
                        style="pointer-events: initial; cursor: pointer"
                        @mousedown="disableMap()"
                        @mouseup="enableMap()"
                      ></v-slider>
                      <v-slider
                        v-else
                        hint="Таймлайн"
                        max="0"
                        min="-50"
                        color="red lighten-2"
                        thumb-color="red"
                        track-color="grey lighten-2"
                        style="pointer-events: initial; cursor: pointer"
                        @mousedown="disableMap()"
                        @mouseup="enableMap()"
                      ></v-slider>
                    </v-col>
                    <v-col cols="1">
                      <v-btn
                        elevation="2"
                        fab
                        x-small
                        style="
                          pointer-events: initial;
                          cursor: pointer;
                          z-index: 999999;
                        "
                      >
                        <v-icon>mdi-play-circle-outline</v-icon>
                      </v-btn>
                    </v-col>
                  </v-row>
                </div>
              --></div>
            </div>
          </v-card>
        </v-col>
      </v-row>
      <div class="uavKartaContainerMobile" v-if="$isMobile()">
        <div id="uavKarta" ref="uavKartaRef"></div>
      </div>
      <!--<div id="uavOverlayMobile" v-if="$isMobile()">
        <v-btn
          :dark="appDark"
          fab
          small
          @click="
            uavControlsUI.createOrder.isOpen = !uavControlsUI.createOrder.isOpen
          "
        >
          <v-icon v-if="uavControlsUI.createOrder.isOpen"> mdi-close </v-icon>
          <v-icon v-else> mdi-plus </v-icon>
        </v-btn>
        <v-snackbar
          class="uav-snackbar"
          v-model="uavControlsUI.createOrder.isOpen"
          vertical
          timeout="-1"
        >
          <h3>Создание заказа</h3>
          <template v-slot:action="{ attrs }">
            <v-btn
              color="indigo"
              text
              v-bind="attrs"
              @click="
                uavControlsUI.createOrder.isOpen =
                  !uavControlsUI.createOrder.isOpen
              "
            >
              Close
            </v-btn>
          </template>
        </v-snackbar>
      </div>-->
    </v-app>
  </div>
</template>

<script>
import "leaflet/dist/leaflet.css";
import L from "leaflet";
import "leaflet-rotatedmarker/leaflet.rotatedMarker.js";
import "leaflet.tilelayer.colorfilter/src/leaflet-tilelayer-colorfilter.min.js";
import axios from "axios";
import "leaflet-pixi-overlay";
import * as PIXI from "pixi.js";
import UI_Route from "@/components/UI_Route";

//let points = [[51.5, -0.09], [60, 45], [50, 3]];
//var icon = L.icon({iconUrl: 'airplane.png', shadowUrl: 'airplane.png', iconSize: [30, 30], shadowSize: [30, 30], iconAnchor: [15, 15], shadowAnchor: [15, 15], popupAnchor: [15, 15]});

export default {
  components: { UI_Route },
  data() {
    return {
      appLoaded: false,
      appDark: true,

      connectionSpeed: "",

      statusBar: {
        connectionSpeed: "—",
        totalUavs: "—",
        totalOrders: "—",
      },
      totalStats: {
        aircraftsInFlight: 0,
        totalOrders: 0,
      },

      length: 2,
      window: 1,

      UI: {
        uavInfo: {
          type: "",
          id: "",
          image: "",
          show: true,
          ordersOnboard: [],
        },
        watchOrder: {
          success: false,
        },
        addOrder: {
          error: false,

          showMarkers: false,
          processing: false,

          locBtnText: "Откуда, куда",
          locBtnTextDefault: "Откуда, куда",
          sending: false,

          arePointsAdded: false,

          firstProcessing: false,
          firstLatLng: null,
          firstMarker: null,
          firstIcon: L.icon({
            iconUrl: "marker-a.png",
            shadowUrl: "marker-a.png",
            iconSize: [30, 30], // size of the icon
            shadowSize: [30, 30], // size of the shadow
            iconAnchor: [15, 15], // point of the icon which will correspond to marker's location
            shadowAnchor: [15, 15], // the same for the shadow
            popupAnchor: [15, 15], // point from which the popup should open relative to the iconAnchor
          }),

          secondProcessing: false,
          secondLatLng: null,
          secondMarker: null,
          secondIcon: L.icon({
            iconUrl: "marker-b.png",
            shadowUrl: "marker-b.png",
            iconSize: [30, 30], // size of the icon
            shadowSize: [30, 30], // size of the shadow
            iconAnchor: [15, 15], // point of the icon which will correspond to marker's location
            shadowAnchor: [15, 15], // the same for the shadow
            popupAnchor: [15, 15], // point from which the popup should open relative to the iconAnchor
          }),

          orderWeight: null,

          orderInfo: {},
        },
        trackOrder: {
          input_id: null
        },
        track: {
          RoutePolyline: Object,
          RoutePoints: [],
          PreRoutePolyline: Object,
          PreRoutePoints: [],
          ApRoutePolyline: Object,
          ApRoutePoints: [],
          result: undefined
        },
      },

      infoMessage: "",

      map: {},
      mapTile: {},
      mapFilter: {
        light: ["bright:80%"],
        dark: [
          "invert:100%",
          "grayscale:65%",
          "bright:80%",
          "contrast:97%",
          "hue:200deg",
          "saturate:150%",
        ],
      },

      aircraftFromApi: {
        raw: null,
      },
      HubsFromApi: {
        small: [],
        medium: [],
        large: [],
      },

      MarkersArrays: [[], [], [], []],
      MarkersArrays_len: [3000, 3000, 2000, 100],
      initPos: 1000.0,

      uavControlsUI: {
        createOrder: {
          isOpen: false,
        },
      },

      apiServer: {
        url: "http://roknyazev.engineer/",
      },

      postHeaders: {
        "Content-Type": "application/json",
        "Content-Length": null,
        Accept: "*/*",
        "Accept-Encoding": "gzip, deflate, br",
        Connection: "keep-alive",
      },

      renderer: null,
      project: null,
      container: null,

      hub_renderer: null,
      hub_project: null,
      hub_container: null,

      DEV_autoRenderUavs: false,

      prev_counters: [0, 0, 0, 0],
      interval: 35,

      getScale: null,
      target_scale: 1 / 7,
      selected_id: 0,
      dist20kmProject: null,

      globalCircle: null,
    };
  },

  mounted() {
    let vh = window.innerHeight * 0.01;
    document.documentElement.style.setProperty("--vh", `${vh}px`);
    window.addEventListener("resize", () => {
      let vh = window.innerHeight * 0.01;
      document.documentElement.style.setProperty("--vh", `${vh}px`);
    });

    this.map = L.map(this.$refs["uavKartaRef"], {
      preferCanvas: true,
      zoomControl: false,
    }).setView([56.010569, 92.852572], 7);

    if (this.appDark) {
      this.mapTile = L.tileLayer
        .colorFilter("https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png", {
          attribution:
            '© <a href="https://www.openstreetmap.org/copyright">OpenStreetMap</a>',
          filter: this.mapFilter.dark,
        })
        .addTo(this.map);
    } else {
      this.mapTile = L.tileLayer
        .colorFilter("https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png", {
          attribution:
            '© <a href="https://www.openstreetmap.org/copyright">OpenStreetMap</a>',
          filter: this.mapFilter.light,
        })
        .addTo(this.map);
    }

    this.UI_selectUav(null, null);
    this.map.doubleClickZoom.disable();
    this.map.on("contextmenu", this.handleAddMarker);
    this.map.on("viewreset", this.updateHubs()); // !!! toggleDeliveryZone !!!

    this.map.on("dblclick", () => {
      this.selected_id = 0;
      this.UI_selectUav(null, null);

      this.map.removeLayer(this.UI.track.RoutePolyline);
      this.map.removeLayer(this.UI.track.PreRoutePolyline);
      this.map.removeLayer(this.UI.track.ApRoutePolyline);
      this.map.removeLayer(this.UI.addOrder.firstMarker)
      this.map.removeLayer(this.UI.addOrder.secondMarker)

      this.UI.uavInfo.ordersOnboard = []
    });
    this.init();

    let point1 = this.project([55.937325, 92.7366]);
    let point2 = this.project([55.921384, 93.057621]);

    this.dist20kmProject = Math.sqrt(
      Math.pow(point1.x - point2.x, 2) + Math.pow(point1.y - point2.y, 2)
    );
    console.log(this.dist20kmProject);

    this.map.createPane("track");
    this.map.getPane("track").style.zIndex = 400;
    
   

    setInterval(
      function run() {
        this.getUavs();
      }.bind(this),
      this.interval
    );

    setInterval(
      function () {
        if (!this.appLoaded) {
          this.init();
        }
      }.bind(this),
      1000
    );

    setInterval(
      function () {
        if (this.appLoaded) {
          this.statusBar.connectionSpeed = this.connectionSpeed;
          this.statusBar.totalUavs = this.totalStats.aircraftsInFlight;
          this.statusBar.totalOrders = this.totalStats.totalOrders;
        }
      }.bind(this),
      1000
    );
  },

  methods: {
    changeMapTheme() {
      this.map.removeLayer(this.mapTile);
      if (this.appDark) {
        this.mapTile = L.tileLayer
          .colorFilter("https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png", {
            attribution:
              '© <a href="https://www.openstreetmap.org/copyright">OpenStreetMap</a>',
            filter: this.mapFilter.dark,
          })
          .addTo(this.map);
      } else {
        this.mapTile = L.tileLayer
          .colorFilter("https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png", {
            attribution:
              '© <a href="https://www.openstreetmap.org/copyright">OpenStreetMap</a>',
            filter: this.mapFilter.light,
          })
          .addTo(this.map);
      }
    },

    return_is_UI_AddOrder_Proc() {
      return this.UI.addOrder.processing ? true : false;
    },

    sendNewOrder() {
      if (
        this.UI.addOrder.firstMarker != null &&
        this.UI.addOrder.secondMarker != null &&
        this.UI.addOrder.orderWeight != null
      ) {
        this.UI.addOrder.error = false;
        this.UI.addOrder.sending = true;

        this.UI.addOrder.firstMarker.dragging.disable();
        this.UI.addOrder.secondMarker.dragging.disable();

        axios
          .post(
            this.apiServer.url + "orders/new",
            {
              weight: Number(this.UI.addOrder.orderWeight),
              first_lat: this.UI.addOrder.firstLatLng.lat,
              first_lon: this.UI.addOrder.firstLatLng.lng,
              last_lat: this.UI.addOrder.secondLatLng.lat,
              last_lon: this.UI.addOrder.secondLatLng.lng,
            },
            {
              timeout: 1000,
              headers: {
                "Content-Type": "application/json",
                Accept: "*/*",
              },
            }
          )
          .then((response) => {
            console.log(response.data);
            this.UI.addOrder.orderInfo = response.data;

            this.drawPreRoute(this.UI.addOrder.orderInfo);
            this.drawRoute(this.UI.addOrder.orderInfo.route);

            this.window = 1;
            this.toggleDeliveryZone();
          })
          .catch((error) => {
            console.log("ERROR\n\n", error);
            this.UI.addOrder.sending = false;
            this.UI.addOrder.error = true;
            //alert(JSON.stringify(error));
            this.UI.addOrder.firstMarker.dragging.enable();
            this.UI.addOrder.secondMarker.dragging.enable();
          })
          .finally(() => {
            this.UI.addOrder.sending = false;
          });
      } else {
        console.log("neok");
      }
    },

    resetOrder() {
      this.UI.addOrder.showMarkers = false;
      this.deleteOrderMarkers();
    },

    drawRoute(_points) {
      this.UI.track.RoutePoints = [];
      this.map.removeLayer(this.UI.track.RoutePolyline);
      this.UI.track.RoutePolyline = Object;

      _points.forEach((_point) => {
        this.UI.track.RoutePoints.push([
          (_point.lan_pos * 180) / Math.PI,
          (_point.lon_pos * 180) / Math.PI,
        ]);
      });
      this.UI.track.RoutePolyline = L.polyline(this.UI.track.RoutePoints, {
        color: "#00ffff",
        opacity: 0.5,
        pane: "track",
      });
      this.UI.track.RoutePolyline.addTo(this.map);
    },

    drawPreRoute(_routeData) {
      this.UI.track.PreRoutePoints = [];
      this.UI.track.ApRoutePoints = [];
      this.map.removeLayer(this.UI.track.PreRoutePolyline);
      this.map.removeLayer(this.UI.track.ApRoutePolyline);
      this.UI.track.PreRoutePolyline = Object;
      this.UI.track.ApRoutePolyline = Object;

      this.UI.track.PreRoutePoints.push([
        _routeData.first_lat,
        _routeData.first_lon,
      ]);
      this.UI.track.PreRoutePoints.push([
        (_routeData.route[0].lan_pos * 180) / Math.PI,
        (_routeData.route[0].lon_pos * 180) / Math.PI,
      ]);
      this.UI.track.ApRoutePoints.push([
        _routeData.last_lat,
        _routeData.last_lon,
      ]);
      this.UI.track.ApRoutePoints.push([
        (_routeData.route[_routeData.route.length - 1].lan_pos * 180) / Math.PI,
        (_routeData.route[_routeData.route.length - 1].lon_pos * 180) / Math.PI,
      ]);

      this.UI.track.PreRoutePolyline = L.polyline(
        this.UI.track.PreRoutePoints,
        { color: "#ff00ff", opacity: 0.5, pane: "track" }
      );
      this.UI.track.PreRoutePolyline.addTo(this.map);
      this.UI.track.ApRoutePolyline = L.polyline(this.UI.track.ApRoutePoints, {
        color: "#ff00ff",
        opacity: 0.5,
        pane: "track",
      });
      this.UI.track.ApRoutePolyline.addTo(this.map);
    },

    addOrderMarkers() {
      this.UI.addOrder.processing = true;
      this.UI.addOrder.showMarkers = true;
      this.UI.addOrder.locBtnText = "Укажите точку отправления";
    },

    deleteOrderMarkers() {
      if (!this.UI.addOrder.showMarkers) {
        if (this.UI.addOrder.firstMarker != null) {
          this.map.removeLayer(this.UI.addOrder.firstMarker);
          this.UI.addOrder.firstMarker = null;
        }
        if (this.UI.addOrder.secondMarker != null) {
          this.map.removeLayer(this.UI.addOrder.secondMarker);
          this.UI.addOrder.secondMarker = null;
        }
      }
    },

    handleAddMarker(e) {
      if (this.UI.addOrder.processing) {
        this.UI.addOrder.firstProcessing = !this.UI.addOrder.firstProcessing;
        //alert(this.UI.addOrder.secondProcessing);
        if (this.UI.addOrder.secondProcessing) {
          this.UI.addOrder.secondLatLng = e.latlng;
          this.UI.addOrder.secondMarker = new L.marker(e.latlng, {
            icon: this.UI.addOrder.secondIcon,
            draggable: true,
          }).addTo(this.map);
          this.UI.addOrder.secondMarker.bindTooltip(
            "Конечная точка<br>" +
              this.UI.addOrder.secondLatLng.lat.toFixed(6) +
              ", " +
              this.UI.addOrder.secondLatLng.lng.toFixed(6),
            {
              permanent: true,
              direction: "right",
              className: "marker-tooltip",
            }
          );

          this.UI.addOrder.secondMarker.on(
            "dragstart",
            this.freezeSecondMarkerTooltip
          );
          this.UI.addOrder.secondMarker.on(
            "dragend",
            this.updateSecondMarkerCoords
          );

          this.UI.addOrder.secondProcessing = false;

          this.UI.addOrder.firstProcessing = false;
          this.UI.addOrder.processing = false;

          this.UI.addOrder.processing = false;
          this.UI.addOrder.locBtnText = this.UI.addOrder.locBtnTextDefault;

          this.UI.addOrder.arePointsAdded = true;
          this.$refs.addOrder_Weight.focus();
        } else if (this.UI.addOrder.firstProcessing) {
          this.UI.addOrder.firstLatLng = e.latlng;
          this.UI.addOrder.firstMarker = new L.marker(e.latlng, {
            icon: this.UI.addOrder.firstIcon,
            draggable: true,
          }).addTo(this.map);
          this.UI.addOrder.firstMarker.bindTooltip(
            "Начальная точка<br>" +
              this.UI.addOrder.firstLatLng.lat.toFixed(6) +
              ", " +
              this.UI.addOrder.firstLatLng.lng.toFixed(6),
            {
              permanent: true,
              direction: "right",
              className: "marker-tooltip",
            }
          );

          this.UI.addOrder.firstMarker.on(
            "dragstart",
            this.freezeFirstMarkerTooltip
          );
          this.UI.addOrder.firstMarker.on(
            "dragend",
            this.updateFirstMarkerCoords
          );

          this.UI.addOrder.firstProcessing = false;

          this.UI.addOrder.secondProcessing = true;
          this.UI.addOrder.locBtnText = "Укажите точку прибытия";
        }
      }
    },

    freezeFirstMarkerTooltip() {
      this.UI.addOrder.firstMarker.setTooltipContent("Начальная точка<br>...", {
        permanent: true,
        direction: "centre",
        className: "marker-tooltip",
      });
    },

    updateFirstMarkerCoords() {
      this.UI.addOrder.firstLatLng = this.UI.addOrder.firstMarker.getLatLng();
      this.UI.addOrder.firstMarker.setTooltipContent(
        "Начальная точка<br>" +
          this.UI.addOrder.firstLatLng.lat.toFixed(6) +
          ", " +
          this.UI.addOrder.firstLatLng.lng.toFixed(6),
        {
          permanent: true,
          direction: "centre",
          className: "marker-tooltip",
        }
      );
    },

    freezeSecondMarkerTooltip() {
      this.UI.addOrder.secondMarker.setTooltipContent("Конечная точка<br>...", {
        permanent: true,
        direction: "centre",
        className: "marker-tooltip",
      });
    },

    updateSecondMarkerCoords() {
      this.UI.addOrder.secondLatLng = this.UI.addOrder.secondMarker.getLatLng();
      this.UI.addOrder.secondMarker.setTooltipContent(
        "Конечная точка<br>" +
          this.UI.addOrder.secondLatLng.lat.toFixed(6) +
          ", " +
          this.UI.addOrder.secondLatLng.lng.toFixed(6),
        {
          permanent: true,
          direction: "centre",
          className: "marker-tooltip",
        }
      );
    },

    initPixiOverlay() {
      let pixiContainer = new PIXI.Container();
      let pixiOverlay = L.pixiOverlay((utils) => {
        this.container = utils.getContainer();
        this.renderer = utils.getRenderer();
        this.project = utils.latLngToLayerPoint;
        this.getScale = utils.getScale;
      }, pixiContainer);

      let loader = new PIXI.Loader();
      loader
        .add("marker1", "drone-1.png")
        .add("marker2", "drone0.png")
        .add("marker3", "helicopter.png")
        .add("marker4", "aircraft.png")
        .add("hub3", "helipad_large.png")
        .add("hub2", "helipad_medium.png")
        .add("hub1", "helipad_small.png");



      loader.load((loader, resources) => {
        let i = 0;
        while (i < this.MarkersArrays_len[0]) {
          let marker = new PIXI.Sprite(resources.marker1.texture);
          marker.x = 0;
          marker.y = 0;
          marker.anchor.set(0.5, 0.5);
          marker.interactive = true;
          marker.buttonMode = true;

          marker.on("pointerup", () => {
            if (this.selected_id === marker.zIndex) {
              this.selected_id = 0;
              this.UI_deselectUav(marker.zIndex, 0);
            } else {
              this.selected_id = marker.zIndex;
              this.UI_selectUav(marker.zIndex, 0);
            }
          });
          this.MarkersArrays[0].push(marker);
          i++;
        }
        i = 0;
        while (i < this.MarkersArrays_len[1]) {
          let marker = new PIXI.Sprite(resources.marker2.texture);
          marker.x = 0;
          marker.y = 0;
          marker.anchor.set(0.5, 0.5);
          marker.interactive = true;
          marker.buttonMode = true;

          marker.on("pointerup", () => {
            if (this.selected_id === marker.zIndex) {
              this.selected_id = 0;
              this.UI_deselectUav(marker.zIndex, 1);
            } else {
              this.selected_id = marker.zIndex;
              this.UI_selectUav(marker.zIndex, 1);
            }
          });
          this.MarkersArrays[1].push(marker);
          i++;
        }
        i = 0;
        while (i < this.MarkersArrays_len[2]) {
          let marker = new PIXI.Sprite(resources.marker3.texture);
          marker.x = 0;
          marker.y = 0;
          marker.anchor.set(0.5, 0.5);
          marker.interactive = true;
          marker.buttonMode = true;

          marker.on("pointerup", () => {
            if (this.selected_id === marker.zIndex) {
              this.selected_id = 0;
              this.UI_deselectUav(marker.zIndex, 2);
            } else {
              this.selected_id = marker.zIndex;
              this.UI_selectUav(marker.zIndex, 2);
            }
          });
          this.MarkersArrays[2].push(marker);
          i++;
        }
        i = 0;
        while (i < this.MarkersArrays_len[3]) {
          let marker = new PIXI.Sprite(resources.marker4.texture);
          marker.x = 0;
          marker.y = 0;
          marker.anchor.set(0.5, 0.5);
          marker.interactive = true;
          marker.buttonMode = true;

          marker.on("pointerup", () => {
            if (this.selected_id === marker.zIndex) {
              this.selected_id = 0;
              this.UI_deselectUav(marker.zIndex, 3);
            } else {
              this.selected_id = marker.zIndex;
              this.UI_selectUav(marker.zIndex, 3);
            }
          });
          this.MarkersArrays[3].push(marker);
          i++;
        }

        this.HubsFromApi.small.forEach((item) => {
          let marker = new PIXI.Sprite(resources.hub1.texture);
          let projection = this.project([
            (item.Latitude * 180) / Math.PI,
            (item.Longitude * 180) / Math.PI,
          ]);
          marker.x = projection.x;
          marker.y = projection.y;
          marker.anchor.set(0.5, 0.5);
          marker.scale.set(1 / 30);
          marker.interactive = true;
          marker.buttonMode = true;
          marker.zIndex = item.Id;

          marker.isClicked = false;
          let circle = new PIXI.Graphics();
          circle.beginFill(0x4287f5, 0.1);
          console.log(this.dist20kmProject);
          circle.lineStyle(1, 0x4287f5, 1);
          circle.drawCircle(projection.x, projection.y, this.dist20kmProject);
          circle.endFill();
          circle.zIndex = 0;

          marker.on("pointerup", () => {
            if (!marker.isClicked) {
              marker.isClicked = true;
              this.container.addChild(circle);
            } else {
              marker.isClicked = false;
              this.container.removeChild(circle);
            }
          });
          this.container.addChild(marker);

          item["marker"] = marker;
          item["circle"] = circle;
        });

        this.HubsFromApi.medium.forEach((item) => {
          let marker = new PIXI.Sprite(resources.hub2.texture);
          let projection = this.project([
            (item.Latitude * 180) / Math.PI,
            (item.Longitude * 180) / Math.PI,
          ]);
          marker.x = projection.x;
          marker.y = projection.y;
          marker.anchor.set(0.5, 0.5);
          marker.scale.set(1 / 25);
          marker.interactive = true;
          marker.buttonMode = true;
          marker.zIndex = item.Id;

          marker.isClicked = false;
          let circle = new PIXI.Graphics();
          circle.beginFill(0xdc2eff, 0.1);
          console.log(this.dist20kmProject);
          circle.lineStyle(1, 0xdc2eff, 1);
          circle.drawCircle(projection.x, projection.y, this.dist20kmProject);
          circle.endFill();
          circle.zIndex = 0;

          marker.on("pointerup", () => {
            if (!marker.isClicked) {
              marker.isClicked = true;
              this.container.addChild(circle);
            } else {
              marker.isClicked = false;
              this.container.removeChild(circle);
            }
          });
          this.container.addChild(marker);

          item["marker"] = marker;
          item["circle"] = circle;
        });

        this.HubsFromApi.large.forEach((item) => {
          let marker = new PIXI.Sprite(resources.hub3.texture);
          let projection = this.project([
            (item.Latitude * 180) / Math.PI,
            (item.Longitude * 180) / Math.PI,
          ]);
          marker.x = projection.x;
          marker.y = projection.y;
          marker.anchor.set(0.5, 0.5);
          marker.scale.set(1 / 20);
          marker.interactive = true;
          marker.buttonMode = true;
          marker.zIndex = item.Id;

          marker.isClicked = false;
          let circle = new PIXI.Graphics();
          circle.beginFill(0x2effab, 0.1);
          console.log(this.dist20kmProject);
          circle.lineStyle(1, 0x2effab, 1);
          circle.drawCircle(projection.x, projection.y, this.dist20kmProject);
          circle.endFill();
          circle.zIndex = 0;

          marker.on("pointerup", () => {
            if (!marker.isClicked) {
              marker.isClicked = true;
              this.container.addChild(circle);
            } else {
              marker.isClicked = false;
              this.container.removeChild(circle);
            }
          });
          this.container.addChild(marker);

          item["marker"] = marker;
          item["circle"] = circle;
        });
      });

      pixiOverlay.addTo(this.map);
    },

    getUavs() {
      let startTime = new Date();
      axios
        .get(this.apiServer.url + "state/binary", {
          timeout: 2000,
          headers: {},
          responseType: "arraybuffer",
        })
        .then((response) => {
          let endTime = new Date();
          this.connectionSpeed = endTime - startTime;
          this.aircraftFromApi.raw = response.data;
          this.updateUavs();
          this.appLoaded = true;
        })
        .catch((error) => {
          console.log(error);
          this.connectionSpeed = ">2000";
        })
        .finally(() => (this.loading = false));
    },
    updateUavs() {
      let zoom = this.map.getZoom();
      let dataView = new DataView(this.aircraftFromApi.raw);
      let length = dataView.getInt32(0, true);
      let counters = [0, 0, 0, 0];
      let i = 4;
      let id;
      let type;
      let lon;
      let lat;
      let az;
      let mk;
      let projection;
      let target_scale = 1 / this.getScale(zoom);
      let order_count = 0;

      if (this.window == 0) length = 0;

      while (i < length + 4) {
        id = dataView.getBigInt64(i, true);
        type = dataView.getInt32(i + 8, true) + 1;
        if (type === -1) {
          type = 0;
        }
        lon = dataView.getFloat32(i + 12, true);
        lat = dataView.getFloat32(i + 16, true);
        az = dataView.getFloat32(i + 20, true);
        order_count += dataView.getInt32(i + 24, true);

        mk = this.MarkersArrays[type][counters[type]];
        this.container.addChild(mk);
        projection = this.project([lat, lon]);
        mk.x = projection.x;
        mk.y = projection.y;
        mk.angle = az;
        mk.zIndex = id;

        if (id === this.selected_id) {
          mk.scale.set(target_scale);
        } else if (zoom >= 9) mk.scale.set(target_scale / 2.5);
        else mk.scale.set(1 / 2);

        counters[type]++;
        i += 28;
      }

      this.totalStats.aircraftsInFlight =
        counters[0] + counters[1] + counters[2] + counters[3];
      this.totalStats.totalOrders = order_count;
      this.renderer.render(this.container);

      for (let i = 0; i < 4; i++) {
        for (let j = counters[i]; j < this.prev_counters[i]; j++) {
          mk = this.MarkersArrays[i][j];
          this.container.removeChild(mk);
        }
      }
      this.prev_counters = counters;
    },

    init() {
      axios
        .get(this.apiServer.url + "hubs/all", { timeout: 2000, headers: {} })
        .then((response) => {
          let _Hubs = response.data;
          //console.log(_Hubs)
          _Hubs.forEach((_hub) => {
            if (_hub.TypeHub == "0") {
              this.HubsFromApi.small.push(_hub);
            } else if (_hub.TypeHub == "1") {
              this.HubsFromApi.medium.push(_hub);
            } else if (_hub.TypeHub == "2") {
              this.HubsFromApi.large.push(_hub);
            } else {
              console.log("error on parsing hubs: invalid type " + _hub);
            }
          });
          this.updateHubs();
          this.appLoaded = true;
        })
        .catch((error) => {
          console.log(error);
        })
        .finally(() => (this.loading = false));
    },
    updateHubs() {
      this.initPixiOverlay();
    },

    UI_selectUav(_uavId, _uavType) {
      try{
        this.map.removeLayer(this.UI.track.RoutePolyline);
        this.map.removeLayer(this.UI.track.PreRoutePolyline);
        this.map.removeLayer(this.UI.track.ApRoutePolyline);
        this.map.removeLayer(this.UI.addOrder.firstMarker)
        this.map.removeLayer(this.UI.addOrder.secondMarker)
      }
      catch(e){
        console.log(e)
      }

      if (_uavId == null) {
        this.UI.uavInfo.type = "Нет информации";
        this.UI.uavInfo.image = "";
        this.UI.uavInfo.id = "";
        return;
      }
      this.UI.uavInfo.id = _uavId;
      if (_uavType == 0) {
        this.UI.uavInfo.type = "БпЛА для доставки на последней миле";
        this.UI.uavInfo.image = "drone-1.png";
      } else if (_uavType == 1) {
        this.UI.uavInfo.type = "Малый БпЛА";
        this.UI.uavInfo.image = "drone0.png";
      } else if (_uavType == 2) {
        this.UI.uavInfo.type = "Беспилотный вертолет";
        this.UI.uavInfo.image = "helicopter.png";
      } else if (_uavType == 3) {
        this.UI.uavInfo.type = "Грузовой самолет";
        this.UI.uavInfo.image = "aircraft.png";
      }
      let orders;
      axios
        .get(this.apiServer.url + "orders/getOrders", {
          timeout: 2000,
          headers: {},
          responseType: "json",
          params: { id: _uavId },
        })
        .then((response) => {
          orders = response.data;
        })
        .catch((error) => {
          console.log(error);
        })
        .finally(() => {
          this.loading = false;
          console.log(orders);
          this.UI.uavInfo.ordersOnboard = orders;
        });
    },

    UI_deselectUav(_uavId, _uavType) {
      this.UI.uavInfo.type = "Нет информации";
      this.UI.uavInfo.image = "";
      this.UI.uavInfo.id = "";
      this.UI.uavInfo.ordersOnboard = [];
      console.log("deselected", _uavId, _uavType);
    },

    disableMap() {
      this.map.dragging.disable();
      this.map.touchZoom.disable();
      this.map.doubleClickZoom.disable();
      this.map.scrollWheelZoom.disable();
      this.map.boxZoom.disable();
      this.map.keyboard.disable();
      if (this.map.tap) this.map.tap.disable();
      document.getElementById("uavKarta").style.cursor = "default";
    },

    enableMap() {
      this.map.dragging.enable();
      this.map.touchZoom.enable();
      this.map.doubleClickZoom.enable();
      this.map.scrollWheelZoom.enable();
      this.map.boxZoom.enable();
      this.map.keyboard.enable();
      if (this.map.tap) this.map.tap.enable();
      document.getElementById("uavKarta").style.cursor = "grab";
    },

    resetAddOrder() {
      if(this.window == 0) {
        this.map.removeLayer(this.UI.track.RoutePolyline);
        this.map.removeLayer(this.UI.track.PreRoutePolyline);
        this.map.removeLayer(this.UI.track.ApRoutePolyline);

        this.map.removeLayer(this.UI.addOrder.firstMarker)
        this.map.removeLayer(this.UI.addOrder.secondMarker)

        this.UI.addOrder.firstMarker = null
        this.UI.addOrder.firstLatLng = null
        this.UI.addOrder.secondMarker = null
        this.UI.addOrder.secondLatLng = null

        this.UI.addOrder.orderWeight = null

        this.UI.addOrder.error = false,

          this.UI.addOrder.showMarkers = false
          this.UI.addOrder.processing = false

          this.UI.addOrder.locBtnText = "Откуда, куда"
          this.UI.addOrder.locBtnTextDefault = "Откуда, куда"
          this.UI.addOrder.sending = false

          this.UI.addOrder.arePointsAdded = false

        this.UI.addOrder.orderInfo = {}
      }
    },

    toggleDeliveryZone() {
      this.HubsFromApi.small.forEach((item) => {
        if(this.window == 0) {
          this.container.addChild(item.circle);
        }
        else {
          this.container.removeChild(item.circle);
        }
      })
      this.HubsFromApi.medium.forEach((item) => {
        if(this.window == 0) {
          this.container.addChild(item.circle);
        }
        else {
          this.container.removeChild(item.circle);
        }
      })
      this.HubsFromApi.large.forEach((item) => {
        if(this.window == 0) {
          this.container.addChild(item.circle);
        }
        else {
          this.container.removeChild(item.circle);
        }
      })
      //console.log(this.container)
    },

    sendTrackOrder(_OrderId) {
      if(_OrderId != null) {
        axios
            .get(
              this.apiServer.url + "orders/getOrderInfo",
              {
                params: {              
                  id: _OrderId
                },
                timeout: 2000
              }
            )
            .then((response) => {
              console.log(response.data);
              this.UI.addOrder.orderInfo = response.data;

              this.UI.track.result = true
  
              this.drawPreRoute(this.UI.addOrder.orderInfo);
              this.drawRoute(this.UI.addOrder.orderInfo.route);

              this.window = 1;

              this.toggleDeliveryZone();
            })
            .catch((error) => {
              console.log("ERROR\n\n", error);

              this.UI.track.result = false
              
              this.map.removeLayer(this.UI.track.RoutePolyline);
              this.map.removeLayer(this.UI.track.PreRoutePolyline);
              this.map.removeLayer(this.UI.track.ApRoutePolyline);
              this.map.removeLayer(this.UI.addOrder.firstMarker)
              this.map.removeLayer(this.UI.addOrder.secondMarker)
            })
            .finally(() => {

            });
      }
    }

  },
};
</script>

<style>
#app {
  height: 100vh; /* Use vh as a fallback for browsers that do not support Custom Properties */
  height: calc(var(--vh, 1vh) * 98);
}
#statusBar {
  height: 3vh;
  height: calc(var(--vh, 1vh) * 3);
  width: 100%;
}
#uavControlsDesktop {
  max-height: 97vh;
  max-height: calc(var(--vh, 1vh) * 97);
}
#uavControlsDesktopCard {
  border: none;
  border-radius: 0;
  width: 100%;
  height: 60%;
  overflow-y: scroll;
  -ms-overflow-style: none; /* IE and Edge */
  scrollbar-width: none; /* Firefox */
}
#uavControlsDesktopCard::-webkit-scrollbar {
  display: none;
}
#uavInfoDesktopCard {
  border: none;
  border-radius: 0;
  width: 100%;
  height: 40%;
  overflow-y: scroll;
  -ms-overflow-style: none; /* IE and Edge */
  scrollbar-width: none; /* Firefox */
}
#uavInfoDesktopCard::-webkit-scrollbar {
  display: none;
}
.uavKartaContainerDesktop {
  height: 97vh;
  height: calc(var(--vh, 1vh) * 97);
  width: 100%;
}
.uavKartaContainerMobile {
  height: 98vh;
  height: calc(var(--vh, 1vh) * 96);
  width: 100%;
}
#uavTimeline {
  position: absolute;
  bottom: 0;
  z-index: 314159;
  pointer-events: none;
  width: 100%;
}
#uavKarta {
  height: 100%;
  width: 100%;
  background-color: black;
}
#devInfo,
#devControls {
  height: 10vh;
  width: 100vw;
}
.column-left {
  float: left;
  width: 33.33%;
  text-align: left;
  padding-left: 5px;
}
.column-center {
  display: inline-block;
  width: 33.33%;
  text-align: center;
}
.column-right {
  float: right;
  width: 33.33%;
  text-align: right;
  padding-right: 5px;
}
.row:after {
  content: "";
  display: table;
  clear: both;
}
.leaflet-container .leaflet-control-attribution {
  color: white;
  background: none;
  background-color: none;
}
.leaflet-container .leaflet-control-attribution a {
  color: white;
}
.leaflet-container .leaflet-control-attribution a:hover {
  text-decoration: underline;
}
.system-bar-element {
  margin-left: 5px;
  margin-right: 5px;
}
#create .v-speed-dial {
  position: absolute;
}
#create .v-btn--floating {
  position: relative;
}
.leaflet-bar {
  box-shadow: 0 1px 5px rgba(0, 0, 0, 0.65);
  border-radius: 5px;
}
#uavOverlayMobile {
  position: absolute;
  bottom: 10px;
  left: 10px;
  z-index: 1001;
}
.uav-snackbar {
  margin-bottom: 50px;
}
.fade-enter-active,
.fade-leave-active {
  transition: opacity 0.2s;
}
.fade-enter, .fade-leave-to /* .fade-leave-active до версии 2.1.8 */ {
  opacity: 0;
}
#uavRoute {
  overflow-y: scroll;
  -ms-overflow-style: none; /* IE and Edge */
  scrollbar-width: none; /* Firefox */
}
#uavRoute::-webkit-scrollbar {
  display: none;
}
.leaflet-container.crosshair-cursor-enabled {
  cursor: crosshair;
}
.marker-tooltip {
  font-size: 1.25em;
  font-family: "Roboto", sans-serif;
  margin: 0px 0px 10px 20px;
  background: rgba(0, 0, 0, 0.65);
  border: none;
  border-radius: 10px;
  color: white;
  text-shadow: none;
  box-shadow: none;
  outline: none;
}
.leaflet-overlay-pane {
  z-index: 500;
  pointer-events: auto;
}
</style>