(window["webpackJsonp"] = window["webpackJsonp"] || []).push([["main"],{

/***/ "./src/$$_lazy_route_resource lazy recursive":
/*!**********************************************************!*\
  !*** ./src/$$_lazy_route_resource lazy namespace object ***!
  \**********************************************************/
/*! no static exports found */
/***/ (function(module, exports) {

function webpackEmptyAsyncContext(req) {
	// Here Promise.resolve().then() is used instead of new Promise() to prevent
	// uncaught exception popping up in devtools
	return Promise.resolve().then(function() {
		var e = new Error("Cannot find module '" + req + "'");
		e.code = 'MODULE_NOT_FOUND';
		throw e;
	});
}
webpackEmptyAsyncContext.keys = function() { return []; };
webpackEmptyAsyncContext.resolve = webpackEmptyAsyncContext;
module.exports = webpackEmptyAsyncContext;
webpackEmptyAsyncContext.id = "./src/$$_lazy_route_resource lazy recursive";

/***/ }),

/***/ "./src/app/app-routing.module.ts":
/*!***************************************!*\
  !*** ./src/app/app-routing.module.ts ***!
  \***************************************/
/*! exports provided: AppRoutingModule */
/***/ (function(module, __webpack_exports__, __webpack_require__) {

"use strict";
__webpack_require__.r(__webpack_exports__);
/* harmony export (binding) */ __webpack_require__.d(__webpack_exports__, "AppRoutingModule", function() { return AppRoutingModule; });
/* harmony import */ var _angular_core__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! @angular/core */ "./node_modules/@angular/core/fesm5/core.js");
/* harmony import */ var _angular_router__WEBPACK_IMPORTED_MODULE_1__ = __webpack_require__(/*! @angular/router */ "./node_modules/@angular/router/fesm5/router.js");
var __decorate = (undefined && undefined.__decorate) || function (decorators, target, key, desc) {
    var c = arguments.length, r = c < 3 ? target : desc === null ? desc = Object.getOwnPropertyDescriptor(target, key) : desc, d;
    if (typeof Reflect === "object" && typeof Reflect.decorate === "function") r = Reflect.decorate(decorators, target, key, desc);
    else for (var i = decorators.length - 1; i >= 0; i--) if (d = decorators[i]) r = (c < 3 ? d(r) : c > 3 ? d(target, key, r) : d(target, key)) || r;
    return c > 3 && r && Object.defineProperty(target, key, r), r;
};


var routes = [];
var AppRoutingModule = /** @class */ (function () {
    function AppRoutingModule() {
    }
    AppRoutingModule = __decorate([
        Object(_angular_core__WEBPACK_IMPORTED_MODULE_0__["NgModule"])({
            imports: [_angular_router__WEBPACK_IMPORTED_MODULE_1__["RouterModule"].forRoot(routes)],
            exports: [_angular_router__WEBPACK_IMPORTED_MODULE_1__["RouterModule"]]
        })
    ], AppRoutingModule);
    return AppRoutingModule;
}());



/***/ }),

/***/ "./src/app/app.component.html":
/*!************************************!*\
  !*** ./src/app/app.component.html ***!
  \************************************/
/*! no static exports found */
/***/ (function(module, exports) {

module.exports = "<mat-toolbar color=\"primary\">\n    <mat-toolbar-row>\n      <app-controls></app-controls>\n      <span class=\"example-spacer\"></span>\n      <mat-icon class=\"example-icon\">pets</mat-icon>\n    </mat-toolbar-row>\n  </mat-toolbar>\n\n  <div style=\"padding:20px;\">\n    <app-emulator></app-emulator>\n    <app-cpu></app-cpu>\n    <app-cia></app-cia>\n    <app-memory></app-memory>\n  </div>\n\n<router-outlet></router-outlet>\n"

/***/ }),

/***/ "./src/app/app.component.scss":
/*!************************************!*\
  !*** ./src/app/app.component.scss ***!
  \************************************/
/*! no static exports found */
/***/ (function(module, exports) {

module.exports = ".example-icon {\n  padding: 0 14px; }\n\n.example-spacer {\n  flex: 1 1 auto; }\n\napp-emulator, app-cpu {\n  margin-right: 20px; }\n\n/*# sourceMappingURL=data:application/json;base64,eyJ2ZXJzaW9uIjozLCJzb3VyY2VzIjpbIi9hcHAvc3JjL2FwcC9hcHAuY29tcG9uZW50LnNjc3MiXSwibmFtZXMiOltdLCJtYXBwaW5ncyI6IkFBQUE7RUFDSSxlQUFlLEVBQUE7O0FBR2pCO0VBQ0UsY0FBYyxFQUFBOztBQUdqQjtFQUNHLGtCQUFrQixFQUFBIiwiZmlsZSI6InNyYy9hcHAvYXBwLmNvbXBvbmVudC5zY3NzIiwic291cmNlc0NvbnRlbnQiOlsiLmV4YW1wbGUtaWNvbiB7XG4gICAgcGFkZGluZzogMCAxNHB4O1xuICB9XG5cbiAgLmV4YW1wbGUtc3BhY2VyIHtcbiAgICBmbGV4OiAxIDEgYXV0bztcbiAgfVxuXG4gYXBwLWVtdWxhdG9yLCBhcHAtY3B1IHtcbiAgICBtYXJnaW4tcmlnaHQ6IDIwcHg7XG4gIH0iXX0= */"

/***/ }),

/***/ "./src/app/app.component.ts":
/*!**********************************!*\
  !*** ./src/app/app.component.ts ***!
  \**********************************/
/*! exports provided: AppComponent */
/***/ (function(module, __webpack_exports__, __webpack_require__) {

"use strict";
__webpack_require__.r(__webpack_exports__);
/* harmony export (binding) */ __webpack_require__.d(__webpack_exports__, "AppComponent", function() { return AppComponent; });
/* harmony import */ var _angular_core__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! @angular/core */ "./node_modules/@angular/core/fesm5/core.js");
var __decorate = (undefined && undefined.__decorate) || function (decorators, target, key, desc) {
    var c = arguments.length, r = c < 3 ? target : desc === null ? desc = Object.getOwnPropertyDescriptor(target, key) : desc, d;
    if (typeof Reflect === "object" && typeof Reflect.decorate === "function") r = Reflect.decorate(decorators, target, key, desc);
    else for (var i = decorators.length - 1; i >= 0; i--) if (d = decorators[i]) r = (c < 3 ? d(r) : c > 3 ? d(target, key, r) : d(target, key)) || r;
    return c > 3 && r && Object.defineProperty(target, key, r), r;
};

var AppComponent = /** @class */ (function () {
    function AppComponent() {
        this.title = 'WebUI';
    }
    AppComponent = __decorate([
        Object(_angular_core__WEBPACK_IMPORTED_MODULE_0__["Component"])({
            selector: 'app-root',
            template: __webpack_require__(/*! ./app.component.html */ "./src/app/app.component.html"),
            styles: [__webpack_require__(/*! ./app.component.scss */ "./src/app/app.component.scss")]
        })
    ], AppComponent);
    return AppComponent;
}());



/***/ }),

/***/ "./src/app/app.module.ts":
/*!*******************************!*\
  !*** ./src/app/app.module.ts ***!
  \*******************************/
/*! exports provided: AppModule */
/***/ (function(module, __webpack_exports__, __webpack_require__) {

"use strict";
__webpack_require__.r(__webpack_exports__);
/* harmony export (binding) */ __webpack_require__.d(__webpack_exports__, "AppModule", function() { return AppModule; });
/* harmony import */ var _angular_platform_browser__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! @angular/platform-browser */ "./node_modules/@angular/platform-browser/fesm5/platform-browser.js");
/* harmony import */ var _angular_core__WEBPACK_IMPORTED_MODULE_1__ = __webpack_require__(/*! @angular/core */ "./node_modules/@angular/core/fesm5/core.js");
/* harmony import */ var _app_routing_module__WEBPACK_IMPORTED_MODULE_2__ = __webpack_require__(/*! ./app-routing.module */ "./src/app/app-routing.module.ts");
/* harmony import */ var _app_component__WEBPACK_IMPORTED_MODULE_3__ = __webpack_require__(/*! ./app.component */ "./src/app/app.component.ts");
/* harmony import */ var _angular_platform_browser_animations__WEBPACK_IMPORTED_MODULE_4__ = __webpack_require__(/*! @angular/platform-browser/animations */ "./node_modules/@angular/platform-browser/fesm5/animations.js");
/* harmony import */ var _angular_material_toolbar__WEBPACK_IMPORTED_MODULE_5__ = __webpack_require__(/*! @angular/material/toolbar */ "./node_modules/@angular/material/esm5/toolbar.es5.js");
/* harmony import */ var _angular_material_icon__WEBPACK_IMPORTED_MODULE_6__ = __webpack_require__(/*! @angular/material/icon */ "./node_modules/@angular/material/esm5/icon.es5.js");
/* harmony import */ var _angular_material_button__WEBPACK_IMPORTED_MODULE_7__ = __webpack_require__(/*! @angular/material/button */ "./node_modules/@angular/material/esm5/button.es5.js");
/* harmony import */ var _logs_logs_component__WEBPACK_IMPORTED_MODULE_8__ = __webpack_require__(/*! ./logs/logs.component */ "./src/app/logs/logs.component.ts");
/* harmony import */ var _services_websocket_service__WEBPACK_IMPORTED_MODULE_9__ = __webpack_require__(/*! ./services/websocket.service */ "./src/app/services/websocket.service.ts");
/* harmony import */ var _services_telemetry_service__WEBPACK_IMPORTED_MODULE_10__ = __webpack_require__(/*! ./services/telemetry.service */ "./src/app/services/telemetry.service.ts");
/* harmony import */ var _cpu_cpu_component__WEBPACK_IMPORTED_MODULE_11__ = __webpack_require__(/*! ./cpu/cpu.component */ "./src/app/cpu/cpu.component.ts");
/* harmony import */ var _memory_memory_component__WEBPACK_IMPORTED_MODULE_12__ = __webpack_require__(/*! ./memory/memory.component */ "./src/app/memory/memory.component.ts");
/* harmony import */ var _emulator_emulator_component__WEBPACK_IMPORTED_MODULE_13__ = __webpack_require__(/*! ./emulator/emulator.component */ "./src/app/emulator/emulator.component.ts");
/* harmony import */ var _angular_material_table__WEBPACK_IMPORTED_MODULE_14__ = __webpack_require__(/*! @angular/material/table */ "./node_modules/@angular/material/esm5/table.es5.js");
/* harmony import */ var _controls_controls_component__WEBPACK_IMPORTED_MODULE_15__ = __webpack_require__(/*! ./controls/controls.component */ "./src/app/controls/controls.component.ts");
/* harmony import */ var _angular_material_grid_list__WEBPACK_IMPORTED_MODULE_16__ = __webpack_require__(/*! @angular/material/grid-list */ "./node_modules/@angular/material/esm5/grid-list.es5.js");
/* harmony import */ var _angular_material_card__WEBPACK_IMPORTED_MODULE_17__ = __webpack_require__(/*! @angular/material/card */ "./node_modules/@angular/material/esm5/card.es5.js");
/* harmony import */ var _cia_cia_component__WEBPACK_IMPORTED_MODULE_18__ = __webpack_require__(/*! ./cia/cia.component */ "./src/app/cia/cia.component.ts");
var __decorate = (undefined && undefined.__decorate) || function (decorators, target, key, desc) {
    var c = arguments.length, r = c < 3 ? target : desc === null ? desc = Object.getOwnPropertyDescriptor(target, key) : desc, d;
    if (typeof Reflect === "object" && typeof Reflect.decorate === "function") r = Reflect.decorate(decorators, target, key, desc);
    else for (var i = decorators.length - 1; i >= 0; i--) if (d = decorators[i]) r = (c < 3 ? d(r) : c > 3 ? d(target, key, r) : d(target, key)) || r;
    return c > 3 && r && Object.defineProperty(target, key, r), r;
};



















var AppModule = /** @class */ (function () {
    function AppModule() {
    }
    AppModule = __decorate([
        Object(_angular_core__WEBPACK_IMPORTED_MODULE_1__["NgModule"])({
            declarations: [
                _app_component__WEBPACK_IMPORTED_MODULE_3__["AppComponent"],
                _logs_logs_component__WEBPACK_IMPORTED_MODULE_8__["LogsComponent"],
                _cpu_cpu_component__WEBPACK_IMPORTED_MODULE_11__["CPUComponent"],
                _memory_memory_component__WEBPACK_IMPORTED_MODULE_12__["MemoryComponent"],
                _emulator_emulator_component__WEBPACK_IMPORTED_MODULE_13__["EmulatorComponent"],
                _controls_controls_component__WEBPACK_IMPORTED_MODULE_15__["ControlsComponent"],
                _cia_cia_component__WEBPACK_IMPORTED_MODULE_18__["CiaComponent"]
            ],
            imports: [
                _angular_platform_browser__WEBPACK_IMPORTED_MODULE_0__["BrowserModule"],
                _app_routing_module__WEBPACK_IMPORTED_MODULE_2__["AppRoutingModule"],
                _angular_platform_browser_animations__WEBPACK_IMPORTED_MODULE_4__["BrowserAnimationsModule"],
                _angular_material_toolbar__WEBPACK_IMPORTED_MODULE_5__["MatToolbarModule"],
                _angular_material_icon__WEBPACK_IMPORTED_MODULE_6__["MatIconModule"],
                _angular_material_button__WEBPACK_IMPORTED_MODULE_7__["MatButtonModule"],
                _angular_material_table__WEBPACK_IMPORTED_MODULE_14__["MatTableModule"],
                _angular_material_grid_list__WEBPACK_IMPORTED_MODULE_16__["MatGridListModule"],
                _angular_material_card__WEBPACK_IMPORTED_MODULE_17__["MatCardModule"]
            ],
            providers: [_services_websocket_service__WEBPACK_IMPORTED_MODULE_9__["WebsocketService"], _services_telemetry_service__WEBPACK_IMPORTED_MODULE_10__["TelemetryService"]],
            bootstrap: [_app_component__WEBPACK_IMPORTED_MODULE_3__["AppComponent"]]
        })
    ], AppModule);
    return AppModule;
}());



/***/ }),

/***/ "./src/app/cia/cia.component.html":
/*!****************************************!*\
  !*** ./src/app/cia/cia.component.html ***!
  \****************************************/
/*! no static exports found */
/***/ (function(module, exports) {

module.exports = "<mat-card class=\"example-card\">\n  <mat-card-header>\n      <mat-card-title>CIA</mat-card-title>\n      <mat-card-subtitle>Chip states</mat-card-subtitle>\n  </mat-card-header>\n\n  <table mat-table [dataSource]=\"dataSource\" class=\"mat-elevation-z8\">\n      <ng-container matColumnDef=\"name\">\n          <th mat-header-cell *matHeaderCellDef> CIA Property </th>\n          <td mat-cell *matCellDef=\"let element\"> {{element.Name}} </td>\n      </ng-container>\n\n      <ng-container matColumnDef=\"value\">\n          <th mat-header-cell *matHeaderCellDef> Value </th>\n          <td mat-cell *matCellDef=\"let element\"> {{element.Value}} </td>\n      </ng-container>\n\n      <tr mat-header-row *matHeaderRowDef=\"displayedColumns\"></tr>\n      <tr mat-row *matRowDef=\"let row; columns: displayedColumns;\"></tr>\n  </table>\n\n  <mat-card-content>\n      <p>\n\n      </p>\n  </mat-card-content>\n</mat-card>"

/***/ }),

/***/ "./src/app/cia/cia.component.scss":
/*!****************************************!*\
  !*** ./src/app/cia/cia.component.scss ***!
  \****************************************/
/*! no static exports found */
/***/ (function(module, exports) {

module.exports = "table {\n  width: 100%; }\n\n.mat-header-row {\n  display: none; }\n\nmat-card {\n  width: 320px;\n  display: inline-block; }\n\n/*# sourceMappingURL=data:application/json;base64,eyJ2ZXJzaW9uIjozLCJzb3VyY2VzIjpbIi9hcHAvc3JjL2FwcC9jaWEvY2lhLmNvbXBvbmVudC5zY3NzIl0sIm5hbWVzIjpbXSwibWFwcGluZ3MiOiJBQUFBO0VBQ0ksV0FBVyxFQUFBOztBQUViO0VBQWlCLGFBQVksRUFBQTs7QUFFN0I7RUFDRSxZQUFZO0VBQ1oscUJBQXFCLEVBQUEiLCJmaWxlIjoic3JjL2FwcC9jaWEvY2lhLmNvbXBvbmVudC5zY3NzIiwic291cmNlc0NvbnRlbnQiOlsidGFibGUge1xuICAgIHdpZHRoOiAxMDAlO1xuICB9XG4gIC5tYXQtaGVhZGVyLXJvdyB7ZGlzcGxheTpub25lO31cblxuICBtYXQtY2FyZCB7XG4gICAgd2lkdGg6IDMyMHB4O1xuICAgIGRpc3BsYXk6IGlubGluZS1ibG9jaztcbiAgfSJdfQ== */"

/***/ }),

/***/ "./src/app/cia/cia.component.ts":
/*!**************************************!*\
  !*** ./src/app/cia/cia.component.ts ***!
  \**************************************/
/*! exports provided: CiaComponent */
/***/ (function(module, __webpack_exports__, __webpack_require__) {

"use strict";
__webpack_require__.r(__webpack_exports__);
/* harmony export (binding) */ __webpack_require__.d(__webpack_exports__, "CiaComponent", function() { return CiaComponent; });
/* harmony import */ var _angular_core__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! @angular/core */ "./node_modules/@angular/core/fesm5/core.js");
/* harmony import */ var _services_telemetry_service__WEBPACK_IMPORTED_MODULE_1__ = __webpack_require__(/*! ../services/telemetry.service */ "./src/app/services/telemetry.service.ts");
/* harmony import */ var _services_tablehelper_service__WEBPACK_IMPORTED_MODULE_2__ = __webpack_require__(/*! ../services/tablehelper.service */ "./src/app/services/tablehelper.service.ts");
var __decorate = (undefined && undefined.__decorate) || function (decorators, target, key, desc) {
    var c = arguments.length, r = c < 3 ? target : desc === null ? desc = Object.getOwnPropertyDescriptor(target, key) : desc, d;
    if (typeof Reflect === "object" && typeof Reflect.decorate === "function") r = Reflect.decorate(decorators, target, key, desc);
    else for (var i = decorators.length - 1; i >= 0; i--) if (d = decorators[i]) r = (c < 3 ? d(r) : c > 3 ? d(target, key, r) : d(target, key)) || r;
    return c > 3 && r && Object.defineProperty(target, key, r), r;
};
var __metadata = (undefined && undefined.__metadata) || function (k, v) {
    if (typeof Reflect === "object" && typeof Reflect.metadata === "function") return Reflect.metadata(k, v);
};



var CiaComponent = /** @class */ (function () {
    function CiaComponent(telemetryService, helper) {
        var _this = this;
        this.telemetryService = telemetryService;
        this.helper = helper;
        this.displayedColumns = ['name', 'value'];
        telemetryService.getTelemetry().subscribe(function (t) {
            var telemetry = JSON.parse(t);
            if (telemetry.Command === 'GetCIAState') {
                _this.dataSource = helper.convertToTableRowsWithBlackList(JSON.parse(atob(telemetry.Payload)), ['Keyboard_matrix']);
            }
        });
    }
    CiaComponent.prototype.ngOnInit = function () {
    };
    CiaComponent = __decorate([
        Object(_angular_core__WEBPACK_IMPORTED_MODULE_0__["Component"])({
            selector: 'app-cia',
            template: __webpack_require__(/*! ./cia.component.html */ "./src/app/cia/cia.component.html"),
            styles: [__webpack_require__(/*! ./cia.component.scss */ "./src/app/cia/cia.component.scss")]
        }),
        __metadata("design:paramtypes", [_services_telemetry_service__WEBPACK_IMPORTED_MODULE_1__["TelemetryService"], _services_tablehelper_service__WEBPACK_IMPORTED_MODULE_2__["TablehelperService"]])
    ], CiaComponent);
    return CiaComponent;
}());



/***/ }),

/***/ "./src/app/controls/controls.component.html":
/*!**************************************************!*\
  !*** ./src/app/controls/controls.component.html ***!
  \**************************************************/
/*! no static exports found */
/***/ (function(module, exports) {

module.exports = "<p>\n  <button *ngIf=\"state.PauseFlag\" (click)=\"startEmulator()\" title=\"Start emulator\" mat-button ><mat-icon class=\"build\">play_arrow</mat-icon></button>\n  <button *ngIf=\"!state.PauseFlag\" (click)=\"stopEmulator()\" title=\"Pause emulator\" mat-button ><mat-icon class=\"build\">pause_arrow</mat-icon></button>\n  <button (click)=\"executeNext()\" title=\"Execute next\" mat-button ><mat-icon class=\"build\">skip_next</mat-icon></button>\n  <button (click)=\"getCPUState()\" title=\"Get CPU state\" mat-button ><mat-icon class=\"build\">memory</mat-icon></button>\n  <button (click)=\"getEmulatorState()\" title=\"Get Emulator state\" mat-button ><mat-icon class=\"build\">computer</mat-icon></button>\n  <button (click)=\"getCIAState()\" title=\"Get CIA state\" mat-button ><mat-icon class=\"build\">keyboard</mat-icon></button>\n  <button (click)=\"getMemoryContent()\" title=\"Get memory content\" mat-button ><mat-icon class=\"build\">sd_card</mat-icon></button>\n  <!-- <button (click)=\"reconnect()\" title=\"Reconnect\" mat-button ><mat-icon class=\"build\">cached</mat-icon></button> -->\n\n  <input #box>\n  <button (click)=\"setBreakpoint(box.value)\" mat-button >Set breakpoint</button>\n\n</p>"

/***/ }),

/***/ "./src/app/controls/controls.component.scss":
/*!**************************************************!*\
  !*** ./src/app/controls/controls.component.scss ***!
  \**************************************************/
/*! no static exports found */
/***/ (function(module, exports) {

module.exports = "\n/*# sourceMappingURL=data:application/json;base64,eyJ2ZXJzaW9uIjozLCJzb3VyY2VzIjpbXSwibmFtZXMiOltdLCJtYXBwaW5ncyI6IiIsImZpbGUiOiJzcmMvYXBwL2NvbnRyb2xzL2NvbnRyb2xzLmNvbXBvbmVudC5zY3NzIn0= */"

/***/ }),

/***/ "./src/app/controls/controls.component.ts":
/*!************************************************!*\
  !*** ./src/app/controls/controls.component.ts ***!
  \************************************************/
/*! exports provided: ControlsComponent */
/***/ (function(module, __webpack_exports__, __webpack_require__) {

"use strict";
__webpack_require__.r(__webpack_exports__);
/* harmony export (binding) */ __webpack_require__.d(__webpack_exports__, "ControlsComponent", function() { return ControlsComponent; });
/* harmony import */ var _angular_core__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! @angular/core */ "./node_modules/@angular/core/fesm5/core.js");
/* harmony import */ var _models_emulatorstate_model__WEBPACK_IMPORTED_MODULE_1__ = __webpack_require__(/*! ../models/emulatorstate.model */ "./src/app/models/emulatorstate.model.ts");
/* harmony import */ var _services_telemetry_service__WEBPACK_IMPORTED_MODULE_2__ = __webpack_require__(/*! ../services/telemetry.service */ "./src/app/services/telemetry.service.ts");
/* harmony import */ var _models_telemetry_model__WEBPACK_IMPORTED_MODULE_3__ = __webpack_require__(/*! ../models/telemetry.model */ "./src/app/models/telemetry.model.ts");
var __decorate = (undefined && undefined.__decorate) || function (decorators, target, key, desc) {
    var c = arguments.length, r = c < 3 ? target : desc === null ? desc = Object.getOwnPropertyDescriptor(target, key) : desc, d;
    if (typeof Reflect === "object" && typeof Reflect.decorate === "function") r = Reflect.decorate(decorators, target, key, desc);
    else for (var i = decorators.length - 1; i >= 0; i--) if (d = decorators[i]) r = (c < 3 ? d(r) : c > 3 ? d(target, key, r) : d(target, key)) || r;
    return c > 3 && r && Object.defineProperty(target, key, r), r;
};
var __metadata = (undefined && undefined.__metadata) || function (k, v) {
    if (typeof Reflect === "object" && typeof Reflect.metadata === "function") return Reflect.metadata(k, v);
};




var ControlsComponent = /** @class */ (function () {
    function ControlsComponent(telemetryService) {
        var _this = this;
        this.telemetryService = telemetryService;
        this.state = new _models_emulatorstate_model__WEBPACK_IMPORTED_MODULE_1__["EmulatorState"]();
        telemetryService.getTelemetry().subscribe(function (t) {
            var telemetry = JSON.parse(t);
            if (telemetry.Command === 'GetEmulatorState') {
                _this.state = JSON.parse(atob(telemetry.Payload));
            }
        });
    }
    ControlsComponent.prototype.ngOnInit = function () {
    };
    ControlsComponent.prototype.setBreakpoint = function (value) {
        var telemetryRequest = new _models_telemetry_model__WEBPACK_IMPORTED_MODULE_3__["Telemetry"]();
        telemetryRequest.Command = 'SetBreakpoint';
        // convert from hex to dec
        var dec = parseInt(value, 16);
        // set the decimal number as a parameter on the command
        telemetryRequest.Parameter = dec.toString();
        this.telemetryService.sendCommand(telemetryRequest);
    };
    ControlsComponent.prototype.startEmulator = function () {
        console.log('start emulator called');
        this.telemetryService.sendStringCommand('Start');
        this.refreshAll();
    };
    ControlsComponent.prototype.stopEmulator = function () {
        console.log('stop emulator called');
        this.telemetryService.sendStringCommand('Stop');
        this.refreshAll();
    };
    ControlsComponent.prototype.executeNext = function () {
        console.log('execute next instruction called');
        this.telemetryService.sendStringCommand('ExecuteNext');
        this.refreshAll();
    };
    ControlsComponent.prototype.getCPUState = function () {
        console.log('get CPU state called');
        this.telemetryService.sendStringCommand('GetCPUState');
    };
    ControlsComponent.prototype.getCIAState = function () {
        console.log('get CIA state called');
        this.telemetryService.sendStringCommand('GetCIAState');
    };
    ControlsComponent.prototype.getEmulatorState = function () {
        console.log('get Emulator state called');
        this.telemetryService.sendStringCommand('GetEmulatorState');
    };
    ControlsComponent.prototype.getMemoryContent = function () {
        console.log('get Memory content called');
        this.telemetryService.sendStringCommand('GetMemoryContent');
    };
    ControlsComponent.prototype.refreshAll = function () {
        var _this = this;
        setTimeout(function () {
            _this.telemetryService.sendStringCommand('GetCPUState');
        }, 50);
        setTimeout(function () {
            _this.telemetryService.sendStringCommand('GetCIAState');
        }, 75);
        setTimeout(function () {
            _this.telemetryService.sendStringCommand('GetEmulatorState');
        }, 100);
        setTimeout(function () {
            _this.telemetryService.sendStringCommand('GetMemoryContent');
        }, 125);
    };
    ControlsComponent = __decorate([
        Object(_angular_core__WEBPACK_IMPORTED_MODULE_0__["Component"])({
            selector: 'app-controls',
            template: __webpack_require__(/*! ./controls.component.html */ "./src/app/controls/controls.component.html"),
            styles: [__webpack_require__(/*! ./controls.component.scss */ "./src/app/controls/controls.component.scss")]
        }),
        __metadata("design:paramtypes", [_services_telemetry_service__WEBPACK_IMPORTED_MODULE_2__["TelemetryService"]])
    ], ControlsComponent);
    return ControlsComponent;
}());



/***/ }),

/***/ "./src/app/cpu/cpu.component.html":
/*!****************************************!*\
  !*** ./src/app/cpu/cpu.component.html ***!
  \****************************************/
/*! no static exports found */
/***/ (function(module, exports) {

module.exports = "<mat-card class=\"example-card\">\n    <mat-card-header>\n        <mat-card-title>CPU</mat-card-title>\n        <mat-card-subtitle>Register states</mat-card-subtitle>\n    </mat-card-header>\n\n    <table mat-table [dataSource]=\"dataSource\" class=\"mat-elevation-z8\">\n        <ng-container matColumnDef=\"name\">\n            <th mat-header-cell *matHeaderCellDef> CPU Property </th>\n            <td mat-cell *matCellDef=\"let element\"> {{element.Name}} </td>\n        </ng-container>\n\n        <ng-container matColumnDef=\"value\">\n            <th mat-header-cell *matHeaderCellDef> Value </th>\n            <td mat-cell *matCellDef=\"let element\"> {{element.Value}} </td>\n        </ng-container>\n\n        <tr mat-header-row *matHeaderRowDef=\"displayedColumns\"></tr>\n        <tr mat-row *matRowDef=\"let row; columns: displayedColumns;\"></tr>\n    </table>\n\n    <mat-card-content>\n        <p>\n\n        </p>\n    </mat-card-content>\n</mat-card>"

/***/ }),

/***/ "./src/app/cpu/cpu.component.scss":
/*!****************************************!*\
  !*** ./src/app/cpu/cpu.component.scss ***!
  \****************************************/
/*! no static exports found */
/***/ (function(module, exports) {

module.exports = "table {\n  width: 100%; }\n\n.mat-header-row {\n  display: none; }\n\nmat-card {\n  width: 300px;\n  display: inline-block; }\n\n/*# sourceMappingURL=data:application/json;base64,eyJ2ZXJzaW9uIjozLCJzb3VyY2VzIjpbIi9hcHAvc3JjL2FwcC9jcHUvY3B1LmNvbXBvbmVudC5zY3NzIl0sIm5hbWVzIjpbXSwibWFwcGluZ3MiOiJBQUFBO0VBQ0ksV0FBVyxFQUFBOztBQUViO0VBQWlCLGFBQVksRUFBQTs7QUFFN0I7RUFDRSxZQUFZO0VBQ1oscUJBQXFCLEVBQUEiLCJmaWxlIjoic3JjL2FwcC9jcHUvY3B1LmNvbXBvbmVudC5zY3NzIiwic291cmNlc0NvbnRlbnQiOlsidGFibGUge1xuICAgIHdpZHRoOiAxMDAlO1xuICB9XG4gIC5tYXQtaGVhZGVyLXJvdyB7ZGlzcGxheTpub25lO31cblxuICBtYXQtY2FyZCB7XG4gICAgd2lkdGg6IDMwMHB4O1xuICAgIGRpc3BsYXk6IGlubGluZS1ibG9jaztcbiAgfSJdfQ== */"

/***/ }),

/***/ "./src/app/cpu/cpu.component.ts":
/*!**************************************!*\
  !*** ./src/app/cpu/cpu.component.ts ***!
  \**************************************/
/*! exports provided: CPUComponent */
/***/ (function(module, __webpack_exports__, __webpack_require__) {

"use strict";
__webpack_require__.r(__webpack_exports__);
/* harmony export (binding) */ __webpack_require__.d(__webpack_exports__, "CPUComponent", function() { return CPUComponent; });
/* harmony import */ var _angular_core__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! @angular/core */ "./node_modules/@angular/core/fesm5/core.js");
/* harmony import */ var _services_telemetry_service__WEBPACK_IMPORTED_MODULE_1__ = __webpack_require__(/*! ../services/telemetry.service */ "./src/app/services/telemetry.service.ts");
/* harmony import */ var _services_tablehelper_service__WEBPACK_IMPORTED_MODULE_2__ = __webpack_require__(/*! ../services/tablehelper.service */ "./src/app/services/tablehelper.service.ts");
var __decorate = (undefined && undefined.__decorate) || function (decorators, target, key, desc) {
    var c = arguments.length, r = c < 3 ? target : desc === null ? desc = Object.getOwnPropertyDescriptor(target, key) : desc, d;
    if (typeof Reflect === "object" && typeof Reflect.decorate === "function") r = Reflect.decorate(decorators, target, key, desc);
    else for (var i = decorators.length - 1; i >= 0; i--) if (d = decorators[i]) r = (c < 3 ? d(r) : c > 3 ? d(target, key, r) : d(target, key)) || r;
    return c > 3 && r && Object.defineProperty(target, key, r), r;
};
var __metadata = (undefined && undefined.__metadata) || function (k, v) {
    if (typeof Reflect === "object" && typeof Reflect.metadata === "function") return Reflect.metadata(k, v);
};



var CPUComponent = /** @class */ (function () {
    function CPUComponent(telemetryService, helper) {
        var _this = this;
        this.telemetryService = telemetryService;
        this.helper = helper;
        this.displayedColumns = ['name', 'value'];
        telemetryService.getTelemetry().subscribe(function (t) {
            var telemetry = JSON.parse(t);
            if (telemetry.Command === 'GetCPUState') {
                _this.dataSource = helper.convertToTableRows(JSON.parse(atob(telemetry.Payload)));
            }
        });
    }
    CPUComponent.prototype.ngOnInit = function () {
    };
    CPUComponent = __decorate([
        Object(_angular_core__WEBPACK_IMPORTED_MODULE_0__["Component"])({
            selector: 'app-cpu',
            template: __webpack_require__(/*! ./cpu.component.html */ "./src/app/cpu/cpu.component.html"),
            styles: [__webpack_require__(/*! ./cpu.component.scss */ "./src/app/cpu/cpu.component.scss")]
        }),
        __metadata("design:paramtypes", [_services_telemetry_service__WEBPACK_IMPORTED_MODULE_1__["TelemetryService"], _services_tablehelper_service__WEBPACK_IMPORTED_MODULE_2__["TablehelperService"]])
    ], CPUComponent);
    return CPUComponent;
}());



/***/ }),

/***/ "./src/app/emulator/emulator.component.html":
/*!**************************************************!*\
  !*** ./src/app/emulator/emulator.component.html ***!
  \**************************************************/
/*! no static exports found */
/***/ (function(module, exports) {

module.exports = "\n<mat-card class=\"example-card\">\n    <mat-card-header>\n        <mat-card-title>Emulator</mat-card-title>\n        <mat-card-subtitle>States</mat-card-subtitle>\n    </mat-card-header>\n\n\n    <table mat-table [dataSource]=\"dataSource\" class=\"mat-elevation-z8\">\n        <ng-container matColumnDef=\"name\">\n            <th mat-header-cell *matHeaderCellDef> Emulator State </th>\n            <td mat-cell *matCellDef=\"let element\"> {{element.Name}} </td>\n        </ng-container>\n\n        <ng-container matColumnDef=\"value\">\n            <th mat-header-cell *matHeaderCellDef> Value </th>\n            <td mat-cell *matCellDef=\"let element\"> {{element.Value}} </td>\n        </ng-container>\n\n        <tr mat-header-row *matHeaderRowDef=\"displayedColumns\"></tr>\n        <tr mat-row *matRowDef=\"let row; columns: displayedColumns;\"></tr>\n    </table>\n\n\n    <mat-card-content>\n        <p>\n\n        </p>\n    </mat-card-content>\n</mat-card>\n\n\n\n"

/***/ }),

/***/ "./src/app/emulator/emulator.component.scss":
/*!**************************************************!*\
  !*** ./src/app/emulator/emulator.component.scss ***!
  \**************************************************/
/*! no static exports found */
/***/ (function(module, exports) {

module.exports = "table {\n  width: 100%; }\n\n.mat-header-row {\n  display: none; }\n\nmat-card {\n  width: 300px;\n  display: inline-block; }\n\n/*# sourceMappingURL=data:application/json;base64,eyJ2ZXJzaW9uIjozLCJzb3VyY2VzIjpbIi9hcHAvc3JjL2FwcC9lbXVsYXRvci9lbXVsYXRvci5jb21wb25lbnQuc2NzcyJdLCJuYW1lcyI6W10sIm1hcHBpbmdzIjoiQUFBQTtFQUNJLFdBQVcsRUFBQTs7QUFHYjtFQUFpQixhQUFZLEVBQUE7O0FBRTdCO0VBQ0UsWUFBWTtFQUNaLHFCQUFxQixFQUFBIiwiZmlsZSI6InNyYy9hcHAvZW11bGF0b3IvZW11bGF0b3IuY29tcG9uZW50LnNjc3MiLCJzb3VyY2VzQ29udGVudCI6WyJ0YWJsZSB7XG4gICAgd2lkdGg6IDEwMCU7XG4gIH1cblxuICAubWF0LWhlYWRlci1yb3cge2Rpc3BsYXk6bm9uZTt9XG5cbiAgbWF0LWNhcmQge1xuICAgIHdpZHRoOiAzMDBweDtcbiAgICBkaXNwbGF5OiBpbmxpbmUtYmxvY2s7XG4gIH0iXX0= */"

/***/ }),

/***/ "./src/app/emulator/emulator.component.ts":
/*!************************************************!*\
  !*** ./src/app/emulator/emulator.component.ts ***!
  \************************************************/
/*! exports provided: EmulatorComponent */
/***/ (function(module, __webpack_exports__, __webpack_require__) {

"use strict";
__webpack_require__.r(__webpack_exports__);
/* harmony export (binding) */ __webpack_require__.d(__webpack_exports__, "EmulatorComponent", function() { return EmulatorComponent; });
/* harmony import */ var _angular_core__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! @angular/core */ "./node_modules/@angular/core/fesm5/core.js");
/* harmony import */ var _services_telemetry_service__WEBPACK_IMPORTED_MODULE_1__ = __webpack_require__(/*! ../services/telemetry.service */ "./src/app/services/telemetry.service.ts");
/* harmony import */ var _services_tablehelper_service__WEBPACK_IMPORTED_MODULE_2__ = __webpack_require__(/*! ../services/tablehelper.service */ "./src/app/services/tablehelper.service.ts");
var __decorate = (undefined && undefined.__decorate) || function (decorators, target, key, desc) {
    var c = arguments.length, r = c < 3 ? target : desc === null ? desc = Object.getOwnPropertyDescriptor(target, key) : desc, d;
    if (typeof Reflect === "object" && typeof Reflect.decorate === "function") r = Reflect.decorate(decorators, target, key, desc);
    else for (var i = decorators.length - 1; i >= 0; i--) if (d = decorators[i]) r = (c < 3 ? d(r) : c > 3 ? d(target, key, r) : d(target, key)) || r;
    return c > 3 && r && Object.defineProperty(target, key, r), r;
};
var __metadata = (undefined && undefined.__metadata) || function (k, v) {
    if (typeof Reflect === "object" && typeof Reflect.metadata === "function") return Reflect.metadata(k, v);
};



var EmulatorComponent = /** @class */ (function () {
    function EmulatorComponent(telemetryService, helper) {
        var _this = this;
        this.telemetryService = telemetryService;
        this.helper = helper;
        this.displayedColumns = ['name', 'value'];
        telemetryService.getTelemetry().subscribe(function (t) {
            var telemetry = JSON.parse(t);
            if (telemetry.Command === 'GetEmulatorState') {
                var emuState = JSON.parse(atob(telemetry.Payload));
                emuState.CycleCount += ' '; // HACK: force this variable to be a string to avoid hex conversion
                _this.dataSource = helper.convertToTableRows(emuState);
            }
        });
    }
    EmulatorComponent.prototype.ngOnInit = function () {
    };
    EmulatorComponent = __decorate([
        Object(_angular_core__WEBPACK_IMPORTED_MODULE_0__["Component"])({
            selector: 'app-emulator',
            template: __webpack_require__(/*! ./emulator.component.html */ "./src/app/emulator/emulator.component.html"),
            styles: [__webpack_require__(/*! ./emulator.component.scss */ "./src/app/emulator/emulator.component.scss")]
        }),
        __metadata("design:paramtypes", [_services_telemetry_service__WEBPACK_IMPORTED_MODULE_1__["TelemetryService"], _services_tablehelper_service__WEBPACK_IMPORTED_MODULE_2__["TablehelperService"]])
    ], EmulatorComponent);
    return EmulatorComponent;
}());



/***/ }),

/***/ "./src/app/logs/logs.component.html":
/*!******************************************!*\
  !*** ./src/app/logs/logs.component.html ***!
  \******************************************/
/*! no static exports found */
/***/ (function(module, exports) {

module.exports = "<p>\n    <button (click)=\"startEmulator()\" mat-button >Start</button>\n    <button (click)=\"stopEmulator()\" mat-button >Stop</button>\n    <button (click)=\"executeNext()\" mat-button >Execute Next</button>\n    <button (click)=\"getCPUState()\" mat-button >Get CPU State</button>\n    <button (click)=\"getEmulatorState()\" mat-button >Get Emulator State</button>\n    <button (click)=\"getMemoryContent()\" mat-button >Get Memory Content</button>\n\n    <br/>\n  {{ latestMessage }}\n\n</p>\n"

/***/ }),

/***/ "./src/app/logs/logs.component.scss":
/*!******************************************!*\
  !*** ./src/app/logs/logs.component.scss ***!
  \******************************************/
/*! no static exports found */
/***/ (function(module, exports) {

module.exports = "\n/*# sourceMappingURL=data:application/json;base64,eyJ2ZXJzaW9uIjozLCJzb3VyY2VzIjpbXSwibmFtZXMiOltdLCJtYXBwaW5ncyI6IiIsImZpbGUiOiJzcmMvYXBwL2xvZ3MvbG9ncy5jb21wb25lbnQuc2NzcyJ9 */"

/***/ }),

/***/ "./src/app/logs/logs.component.ts":
/*!****************************************!*\
  !*** ./src/app/logs/logs.component.ts ***!
  \****************************************/
/*! exports provided: LogsComponent */
/***/ (function(module, __webpack_exports__, __webpack_require__) {

"use strict";
__webpack_require__.r(__webpack_exports__);
/* harmony export (binding) */ __webpack_require__.d(__webpack_exports__, "LogsComponent", function() { return LogsComponent; });
/* harmony import */ var _angular_core__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! @angular/core */ "./node_modules/@angular/core/fesm5/core.js");
/* harmony import */ var _services_telemetry_service__WEBPACK_IMPORTED_MODULE_1__ = __webpack_require__(/*! ../services/telemetry.service */ "./src/app/services/telemetry.service.ts");
var __decorate = (undefined && undefined.__decorate) || function (decorators, target, key, desc) {
    var c = arguments.length, r = c < 3 ? target : desc === null ? desc = Object.getOwnPropertyDescriptor(target, key) : desc, d;
    if (typeof Reflect === "object" && typeof Reflect.decorate === "function") r = Reflect.decorate(decorators, target, key, desc);
    else for (var i = decorators.length - 1; i >= 0; i--) if (d = decorators[i]) r = (c < 3 ? d(r) : c > 3 ? d(target, key, r) : d(target, key)) || r;
    return c > 3 && r && Object.defineProperty(target, key, r), r;
};
var __metadata = (undefined && undefined.__metadata) || function (k, v) {
    if (typeof Reflect === "object" && typeof Reflect.metadata === "function") return Reflect.metadata(k, v);
};


var LogsComponent = /** @class */ (function () {
    function LogsComponent(telemetryService) {
        var _this = this;
        this.telemetryService = telemetryService;
        telemetryService.getTelemetry().subscribe(function (t) {
            var telemetry = JSON.parse(t);
            console.log(telemetry);
            console.log(atob(telemetry.Payload));
            // console.log("message received");
            _this.latestMessage = atob(telemetry.Payload);
        });
    }
    LogsComponent.prototype.ngOnInit = function () {
    };
    LogsComponent.prototype.startEmulator = function () {
        console.log('start emulator called');
        this.telemetryService.sendStringCommand('Start');
    };
    LogsComponent.prototype.stopEmulator = function () {
        console.log('stop emulator called');
        this.telemetryService.sendStringCommand('Stop');
    };
    LogsComponent.prototype.executeNext = function () {
        console.log('execute next instruction called');
        this.telemetryService.sendStringCommand('ExecuteNext');
    };
    LogsComponent.prototype.getCPUState = function () {
        console.log('get CPU state called');
        this.telemetryService.sendStringCommand('GetCPUState');
    };
    LogsComponent.prototype.getEmulatorState = function () {
        console.log('get Emulator state called');
        this.telemetryService.sendStringCommand('GetEmulatorState');
    };
    LogsComponent.prototype.getMemoryContent = function () {
        console.log('get Memory content called');
        this.telemetryService.sendStringCommand('GetMemoryContent');
    };
    LogsComponent = __decorate([
        Object(_angular_core__WEBPACK_IMPORTED_MODULE_0__["Component"])({
            selector: 'app-logs',
            template: __webpack_require__(/*! ./logs.component.html */ "./src/app/logs/logs.component.html"),
            styles: [__webpack_require__(/*! ./logs.component.scss */ "./src/app/logs/logs.component.scss")]
        }),
        __metadata("design:paramtypes", [_services_telemetry_service__WEBPACK_IMPORTED_MODULE_1__["TelemetryService"]])
    ], LogsComponent);
    return LogsComponent;
}());



/***/ }),

/***/ "./src/app/memory/memory.component.html":
/*!**********************************************!*\
  !*** ./src/app/memory/memory.component.html ***!
  \**********************************************/
/*! no static exports found */
/***/ (function(module, exports) {

module.exports = "\n<mat-card class=\"example-card\">\n        <mat-card-header>\n            <mat-card-title>Memory</mat-card-title>\n            <mat-card-subtitle>Program counter window</mat-card-subtitle>\n        </mat-card-header>\n\n\n\n\n<mat-card-content #dataContainer>\n        <div ></div>\n</mat-card-content>\n</mat-card>\n"

/***/ }),

/***/ "./src/app/memory/memory.component.scss":
/*!**********************************************!*\
  !*** ./src/app/memory/memory.component.scss ***!
  \**********************************************/
/*! no static exports found */
/***/ (function(module, exports) {

module.exports = "mat-card {\n  width: 400px;\n  display: inline-grid; }\n\n.current-pc-highlight {\n  background-color: yellow;\n  color: red; }\n\n/*# sourceMappingURL=data:application/json;base64,eyJ2ZXJzaW9uIjozLCJzb3VyY2VzIjpbIi9hcHAvc3JjL2FwcC9tZW1vcnkvbWVtb3J5LmNvbXBvbmVudC5zY3NzIl0sIm5hbWVzIjpbXSwibWFwcGluZ3MiOiJBQUFBO0VBQ0ksWUFBWTtFQUNaLG9CQUFvQixFQUFBOztBQUd4QjtFQUNNLHdCQUF3QjtFQUN4QixVQUFVLEVBQUEiLCJmaWxlIjoic3JjL2FwcC9tZW1vcnkvbWVtb3J5LmNvbXBvbmVudC5zY3NzIiwic291cmNlc0NvbnRlbnQiOlsibWF0LWNhcmQge1xuICAgIHdpZHRoOiA0MDBweDtcbiAgICBkaXNwbGF5OiBpbmxpbmUtZ3JpZDtcbiAgfVxuXG4uY3VycmVudC1wYy1oaWdobGlnaHQge1xuICAgICAgYmFja2dyb3VuZC1jb2xvcjogeWVsbG93O1xuICAgICAgY29sb3I6IHJlZDtcbiAgfSJdfQ== */"

/***/ }),

/***/ "./src/app/memory/memory.component.ts":
/*!********************************************!*\
  !*** ./src/app/memory/memory.component.ts ***!
  \********************************************/
/*! exports provided: MemoryComponent */
/***/ (function(module, __webpack_exports__, __webpack_require__) {

"use strict";
__webpack_require__.r(__webpack_exports__);
/* harmony export (binding) */ __webpack_require__.d(__webpack_exports__, "MemoryComponent", function() { return MemoryComponent; });
/* harmony import */ var _angular_core__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! @angular/core */ "./node_modules/@angular/core/fesm5/core.js");
/* harmony import */ var _services_telemetry_service__WEBPACK_IMPORTED_MODULE_1__ = __webpack_require__(/*! ../services/telemetry.service */ "./src/app/services/telemetry.service.ts");
/* harmony import */ var _models_cpustate_model__WEBPACK_IMPORTED_MODULE_2__ = __webpack_require__(/*! ../models/cpustate.model */ "./src/app/models/cpustate.model.ts");
var __decorate = (undefined && undefined.__decorate) || function (decorators, target, key, desc) {
    var c = arguments.length, r = c < 3 ? target : desc === null ? desc = Object.getOwnPropertyDescriptor(target, key) : desc, d;
    if (typeof Reflect === "object" && typeof Reflect.decorate === "function") r = Reflect.decorate(decorators, target, key, desc);
    else for (var i = decorators.length - 1; i >= 0; i--) if (d = decorators[i]) r = (c < 3 ? d(r) : c > 3 ? d(target, key, r) : d(target, key)) || r;
    return c > 3 && r && Object.defineProperty(target, key, r), r;
};
var __metadata = (undefined && undefined.__metadata) || function (k, v) {
    if (typeof Reflect === "object" && typeof Reflect.metadata === "function") return Reflect.metadata(k, v);
};



var MemoryComponent = /** @class */ (function () {
    function MemoryComponent(telemetryService) {
        var _this = this;
        this.telemetryService = telemetryService;
        this.memoryContent = 'blank';
        this.cpuState = new _models_cpustate_model__WEBPACK_IMPORTED_MODULE_2__["CPUState"]();
        telemetryService.getTelemetry().subscribe(function (t) {
            var telemetry = JSON.parse(t);
            if (telemetry.Command === 'GetCPUState') {
                _this.cpuState = JSON.parse(atob(telemetry.Payload));
            }
            if (telemetry.Command === 'GetMemoryContent') {
                var byteCharacters = atob(telemetry.Payload);
                var byteNumbers = new Array(byteCharacters.length);
                for (var i = 0; i < byteCharacters.length; i++) {
                    byteNumbers[i] = byteCharacters.charCodeAt(i);
                }
                var byteArray = new Uint8Array(byteNumbers);
                //console.log(byteArray);
                var hexResult = '';
                var cnt = 0;
                var windowSize = 100;
                var windowStart = 0;
                var windowEnd = 0;
                windowStart = _this.cpuState.PC - windowSize;
                while (windowStart % 16 !== 0) {
                    windowStart = windowStart - 1;
                }
                windowEnd = _this.cpuState.PC + windowSize;
                while (windowEnd % 16 !== 15) {
                    windowEnd = windowEnd + 1;
                }
                for (var _i = 0, byteNumbers_1 = byteNumbers; _i < byteNumbers_1.length; _i++) {
                    var byte = byteNumbers_1[_i];
                    // only print the window
                    if (cnt >= windowStart && cnt <= windowEnd) {
                        // highlight the current PC instruction
                        if (cnt === _this.cpuState.PC) {
                            hexResult += '<span style="background-color: yellow;color: red;" class="current-pc-highlight">';
                        }
                        // tslint:disable-next-line:max-line-length
                        hexResult += (cnt % 16 ? ' ' : '<br/>' + (1e7 + (cnt).toString(16)).slice(-8) + ' | ') + (1e7 + byteArray[cnt].toString(16)).slice(-2);
                        if (cnt === _this.cpuState.PC) {
                            hexResult += '</span>';
                        }
                    }
                    cnt++;
                }
                _this.dataContainer.nativeElement.innerHTML = hexResult;
            }
        });
    }
    MemoryComponent.prototype.loadData = function (data) {
        this.dataContainer.nativeElement.innerHTML = data;
    };
    MemoryComponent.prototype.ngOnInit = function () {
    };
    __decorate([
        Object(_angular_core__WEBPACK_IMPORTED_MODULE_0__["ViewChild"])('dataContainer'),
        __metadata("design:type", _angular_core__WEBPACK_IMPORTED_MODULE_0__["ElementRef"])
    ], MemoryComponent.prototype, "dataContainer", void 0);
    MemoryComponent = __decorate([
        Object(_angular_core__WEBPACK_IMPORTED_MODULE_0__["Component"])({
            selector: 'app-memory',
            template: __webpack_require__(/*! ./memory.component.html */ "./src/app/memory/memory.component.html"),
            styles: [__webpack_require__(/*! ./memory.component.scss */ "./src/app/memory/memory.component.scss")]
        }),
        __metadata("design:paramtypes", [_services_telemetry_service__WEBPACK_IMPORTED_MODULE_1__["TelemetryService"]])
    ], MemoryComponent);
    return MemoryComponent;
}());



/***/ }),

/***/ "./src/app/models/cpustate.model.ts":
/*!******************************************!*\
  !*** ./src/app/models/cpustate.model.ts ***!
  \******************************************/
/*! exports provided: CPUState */
/***/ (function(module, __webpack_exports__, __webpack_require__) {

"use strict";
__webpack_require__.r(__webpack_exports__);
/* harmony export (binding) */ __webpack_require__.d(__webpack_exports__, "CPUState", function() { return CPUState; });
var CPUState = /** @class */ (function () {
    function CPUState() {
    }
    return CPUState;
}());



/***/ }),

/***/ "./src/app/models/emulatorstate.model.ts":
/*!***********************************************!*\
  !*** ./src/app/models/emulatorstate.model.ts ***!
  \***********************************************/
/*! exports provided: EmulatorState */
/***/ (function(module, __webpack_exports__, __webpack_require__) {

"use strict";
__webpack_require__.r(__webpack_exports__);
/* harmony export (binding) */ __webpack_require__.d(__webpack_exports__, "EmulatorState", function() { return EmulatorState; });
var EmulatorState = /** @class */ (function () {
    function EmulatorState() {
    }
    return EmulatorState;
}());



/***/ }),

/***/ "./src/app/models/tablerow.model.ts":
/*!******************************************!*\
  !*** ./src/app/models/tablerow.model.ts ***!
  \******************************************/
/*! exports provided: TableRow */
/***/ (function(module, __webpack_exports__, __webpack_require__) {

"use strict";
__webpack_require__.r(__webpack_exports__);
/* harmony export (binding) */ __webpack_require__.d(__webpack_exports__, "TableRow", function() { return TableRow; });
var TableRow = /** @class */ (function () {
    function TableRow() {
    }
    return TableRow;
}());



/***/ }),

/***/ "./src/app/models/telemetry.model.ts":
/*!*******************************************!*\
  !*** ./src/app/models/telemetry.model.ts ***!
  \*******************************************/
/*! exports provided: Telemetry */
/***/ (function(module, __webpack_exports__, __webpack_require__) {

"use strict";
__webpack_require__.r(__webpack_exports__);
/* harmony export (binding) */ __webpack_require__.d(__webpack_exports__, "Telemetry", function() { return Telemetry; });
var Telemetry = /** @class */ (function () {
    function Telemetry() {
    }
    return Telemetry;
}());



/***/ }),

/***/ "./src/app/services/tablehelper.service.ts":
/*!*************************************************!*\
  !*** ./src/app/services/tablehelper.service.ts ***!
  \*************************************************/
/*! exports provided: TablehelperService */
/***/ (function(module, __webpack_exports__, __webpack_require__) {

"use strict";
__webpack_require__.r(__webpack_exports__);
/* harmony export (binding) */ __webpack_require__.d(__webpack_exports__, "TablehelperService", function() { return TablehelperService; });
/* harmony import */ var _angular_core__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! @angular/core */ "./node_modules/@angular/core/fesm5/core.js");
/* harmony import */ var _models_tablerow_model__WEBPACK_IMPORTED_MODULE_1__ = __webpack_require__(/*! ../models/tablerow.model */ "./src/app/models/tablerow.model.ts");
var __decorate = (undefined && undefined.__decorate) || function (decorators, target, key, desc) {
    var c = arguments.length, r = c < 3 ? target : desc === null ? desc = Object.getOwnPropertyDescriptor(target, key) : desc, d;
    if (typeof Reflect === "object" && typeof Reflect.decorate === "function") r = Reflect.decorate(decorators, target, key, desc);
    else for (var i = decorators.length - 1; i >= 0; i--) if (d = decorators[i]) r = (c < 3 ? d(r) : c > 3 ? d(target, key, r) : d(target, key)) || r;
    return c > 3 && r && Object.defineProperty(target, key, r), r;
};
var __metadata = (undefined && undefined.__metadata) || function (k, v) {
    if (typeof Reflect === "object" && typeof Reflect.metadata === "function") return Reflect.metadata(k, v);
};


var TablehelperService = /** @class */ (function () {
    function TablehelperService() {
    }
    TablehelperService.prototype.convertToTableRows = function (input) {
        return this.convertToTableRowsWithBlackList(input, Array());
    };
    TablehelperService.prototype.convertToTableRowsWithBlackList = function (input, blackList) {
        var result = [];
        // Step 1. Get all the object keys.
        var properties = Object.keys(input);
        // Step 3. Iterate throw all keys.
        for (var _i = 0, properties_1 = properties; _i < properties_1.length; _i++) {
            var prop = properties_1[_i];
            if (!blackList.includes(prop)) {
                var row = new _models_tablerow_model__WEBPACK_IMPORTED_MODULE_1__["TableRow"]();
                row.Name = prop.toString();
                var value = input[prop];
                if (typeof (value) === 'number') {
                    value = value.toString(16);
                }
                row.Value = value;
                result.push(row);
            }
        }
        return result;
    };
    TablehelperService = __decorate([
        Object(_angular_core__WEBPACK_IMPORTED_MODULE_0__["Injectable"])({
            providedIn: 'root'
        }),
        __metadata("design:paramtypes", [])
    ], TablehelperService);
    return TablehelperService;
}());



/***/ }),

/***/ "./src/app/services/telemetry.service.ts":
/*!***********************************************!*\
  !*** ./src/app/services/telemetry.service.ts ***!
  \***********************************************/
/*! exports provided: TelemetryService */
/***/ (function(module, __webpack_exports__, __webpack_require__) {

"use strict";
__webpack_require__.r(__webpack_exports__);
/* harmony export (binding) */ __webpack_require__.d(__webpack_exports__, "TelemetryService", function() { return TelemetryService; });
/* harmony import */ var _angular_core__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! @angular/core */ "./node_modules/@angular/core/fesm5/core.js");
/* harmony import */ var _websocket_service__WEBPACK_IMPORTED_MODULE_1__ = __webpack_require__(/*! ./websocket.service */ "./src/app/services/websocket.service.ts");
/* harmony import */ var rxjs__WEBPACK_IMPORTED_MODULE_2__ = __webpack_require__(/*! rxjs */ "./node_modules/rxjs/_esm5/index.js");
/* harmony import */ var _models_telemetry_model__WEBPACK_IMPORTED_MODULE_3__ = __webpack_require__(/*! ../models/telemetry.model */ "./src/app/models/telemetry.model.ts");
var __decorate = (undefined && undefined.__decorate) || function (decorators, target, key, desc) {
    var c = arguments.length, r = c < 3 ? target : desc === null ? desc = Object.getOwnPropertyDescriptor(target, key) : desc, d;
    if (typeof Reflect === "object" && typeof Reflect.decorate === "function") r = Reflect.decorate(decorators, target, key, desc);
    else for (var i = decorators.length - 1; i >= 0; i--) if (d = decorators[i]) r = (c < 3 ? d(r) : c > 3 ? d(target, key, r) : d(target, key)) || r;
    return c > 3 && r && Object.defineProperty(target, key, r), r;
};
var __metadata = (undefined && undefined.__metadata) || function (k, v) {
    if (typeof Reflect === "object" && typeof Reflect.metadata === "function") return Reflect.metadata(k, v);
};




var TelemetryService = /** @class */ (function () {
    function TelemetryService(webSocket) {
        this.webSocket = webSocket;
        var openSubscriber = rxjs__WEBPACK_IMPORTED_MODULE_2__["Subscriber"].create(function () { return console.log('connection opened'); });
        this.observableSocket = this.webSocket.createObservableSocket('ws://localhost:8080/ws', openSubscriber);
    }
    TelemetryService.prototype.getTelemetry = function () {
        return this.observableSocket;
    };
    TelemetryService.prototype.sendStringCommand = function (command) {
        var telemetryRequest = new _models_telemetry_model__WEBPACK_IMPORTED_MODULE_3__["Telemetry"]();
        telemetryRequest.Command = command;
        this.sendCommand(telemetryRequest);
    };
    TelemetryService.prototype.sendCommand = function (request) {
        this.webSocket.send(JSON.stringify(request));
    };
    TelemetryService = __decorate([
        Object(_angular_core__WEBPACK_IMPORTED_MODULE_0__["Injectable"])({
            providedIn: 'root'
        }),
        __metadata("design:paramtypes", [_websocket_service__WEBPACK_IMPORTED_MODULE_1__["WebsocketService"]])
    ], TelemetryService);
    return TelemetryService;
}());



/***/ }),

/***/ "./src/app/services/websocket.service.ts":
/*!***********************************************!*\
  !*** ./src/app/services/websocket.service.ts ***!
  \***********************************************/
/*! exports provided: WebsocketService */
/***/ (function(module, __webpack_exports__, __webpack_require__) {

"use strict";
__webpack_require__.r(__webpack_exports__);
/* harmony export (binding) */ __webpack_require__.d(__webpack_exports__, "WebsocketService", function() { return WebsocketService; });
/* harmony import */ var _angular_core__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! @angular/core */ "./node_modules/@angular/core/fesm5/core.js");
/* harmony import */ var rxjs__WEBPACK_IMPORTED_MODULE_1__ = __webpack_require__(/*! rxjs */ "./node_modules/rxjs/_esm5/index.js");
var __decorate = (undefined && undefined.__decorate) || function (decorators, target, key, desc) {
    var c = arguments.length, r = c < 3 ? target : desc === null ? desc = Object.getOwnPropertyDescriptor(target, key) : desc, d;
    if (typeof Reflect === "object" && typeof Reflect.decorate === "function") r = Reflect.decorate(decorators, target, key, desc);
    else for (var i = decorators.length - 1; i >= 0; i--) if (d = decorators[i]) r = (c < 3 ? d(r) : c > 3 ? d(target, key, r) : d(target, key)) || r;
    return c > 3 && r && Object.defineProperty(target, key, r), r;
};


var WebsocketService = /** @class */ (function () {
    function WebsocketService() {
        this.subject = new rxjs__WEBPACK_IMPORTED_MODULE_1__["Subject"]();
    }
    WebsocketService.prototype.createObservableSocket = function (url, openSubscriber) {
        var _this = this;
        this.ws = new WebSocket(url);
        new rxjs__WEBPACK_IMPORTED_MODULE_1__["Observable"](function (observer) {
            _this.ws.onmessage = function (event) { return observer.next(event.data); };
            _this.ws.onerror = function (event) { return observer.error(event); };
            _this.ws.onclose = function (event) { return observer.complete(); };
            _this.ws.onopen = function (event) {
                openSubscriber.next();
                openSubscriber.complete();
            };
            return function () { return _this.ws.close(); };
        }).subscribe(function (data) { _this.subject.next(data); console.log(data); });
        return this.subject;
    };
    WebsocketService.prototype.send = function (message) {
        this.ws.send(message);
    };
    WebsocketService = __decorate([
        Object(_angular_core__WEBPACK_IMPORTED_MODULE_0__["Injectable"])({
            providedIn: 'root'
        })
    ], WebsocketService);
    return WebsocketService;
}());



/***/ }),

/***/ "./src/environments/environment.ts":
/*!*****************************************!*\
  !*** ./src/environments/environment.ts ***!
  \*****************************************/
/*! exports provided: environment */
/***/ (function(module, __webpack_exports__, __webpack_require__) {

"use strict";
__webpack_require__.r(__webpack_exports__);
/* harmony export (binding) */ __webpack_require__.d(__webpack_exports__, "environment", function() { return environment; });
// This file can be replaced during build by using the `fileReplacements` array.
// `ng build --prod` replaces `environment.ts` with `environment.prod.ts`.
// The list of file replacements can be found in `angular.json`.
var environment = {
    production: false
};
/*
 * For easier debugging in development mode, you can import the following file
 * to ignore zone related error stack frames such as `zone.run`, `zoneDelegate.invokeTask`.
 *
 * This import should be commented out in production mode because it will have a negative impact
 * on performance if an error is thrown.
 */
// import 'zone.js/dist/zone-error';  // Included with Angular CLI.


/***/ }),

/***/ "./src/main.ts":
/*!*********************!*\
  !*** ./src/main.ts ***!
  \*********************/
/*! no exports provided */
/***/ (function(module, __webpack_exports__, __webpack_require__) {

"use strict";
__webpack_require__.r(__webpack_exports__);
/* harmony import */ var _angular_core__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! @angular/core */ "./node_modules/@angular/core/fesm5/core.js");
/* harmony import */ var _angular_platform_browser_dynamic__WEBPACK_IMPORTED_MODULE_1__ = __webpack_require__(/*! @angular/platform-browser-dynamic */ "./node_modules/@angular/platform-browser-dynamic/fesm5/platform-browser-dynamic.js");
/* harmony import */ var _app_app_module__WEBPACK_IMPORTED_MODULE_2__ = __webpack_require__(/*! ./app/app.module */ "./src/app/app.module.ts");
/* harmony import */ var _environments_environment__WEBPACK_IMPORTED_MODULE_3__ = __webpack_require__(/*! ./environments/environment */ "./src/environments/environment.ts");




if (_environments_environment__WEBPACK_IMPORTED_MODULE_3__["environment"].production) {
    Object(_angular_core__WEBPACK_IMPORTED_MODULE_0__["enableProdMode"])();
}
Object(_angular_platform_browser_dynamic__WEBPACK_IMPORTED_MODULE_1__["platformBrowserDynamic"])().bootstrapModule(_app_app_module__WEBPACK_IMPORTED_MODULE_2__["AppModule"])
    .catch(function (err) { return console.error(err); });


/***/ }),

/***/ 0:
/*!***************************!*\
  !*** multi ./src/main.ts ***!
  \***************************/
/*! no static exports found */
/***/ (function(module, exports, __webpack_require__) {

module.exports = __webpack_require__(/*! /app/src/main.ts */"./src/main.ts");


/***/ })

},[[0,"runtime","vendor"]]]);
//# sourceMappingURL=main.js.map