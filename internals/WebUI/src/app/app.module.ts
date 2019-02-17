import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';

import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import { MatToolbarModule } from '@angular/material/toolbar';
import { MatIconModule } from '@angular/material/icon';
import { MatButtonModule } from '@angular/material/button';
import { LogsComponent } from './logs/logs.component';
import { WebsocketService } from './services/websocket.service';
import { TelemetryService } from './services/telemetry.service';
import { CPUComponent } from './cpu/cpu.component';
import { MemoryComponent } from './memory/memory.component';
import { EmulatorComponent } from './emulator/emulator.component';
import { MatTableModule } from '@angular/material/table';
import { ControlsComponent } from './controls/controls.component';
import {MatGridListModule} from '@angular/material/grid-list';
import {MatCardModule} from '@angular/material/card';


@NgModule({
  declarations: [
    AppComponent,
    LogsComponent,
    CPUComponent,
    MemoryComponent,
    EmulatorComponent,
    ControlsComponent
  ],
  imports: [
    BrowserModule,
    AppRoutingModule,
    BrowserAnimationsModule,
    MatToolbarModule,
    MatIconModule,
    MatButtonModule,
    MatTableModule,
    MatGridListModule,
    MatCardModule
  ],
  providers: [WebsocketService, TelemetryService],
  bootstrap: [AppComponent]
})
export class AppModule { }
