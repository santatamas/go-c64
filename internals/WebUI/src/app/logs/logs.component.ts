import { Component, OnInit } from '@angular/core';
import { TelemetryService } from '../services/telemetry.service';
import { Telemetry } from '../models/telemetry.model';

@Component({
  selector: 'app-logs',
  templateUrl: './logs.component.html',
  styleUrls: ['./logs.component.scss']
})
export class LogsComponent implements OnInit {

  latestMessage: string;

  constructor(private telemetryService: TelemetryService) {
    telemetryService.getTelemetry().subscribe((t: string) => {
      const telemetry: Telemetry = JSON.parse(t);
      console.log(telemetry);
      console.log(atob(telemetry.Payload));
      // console.log("message received");
      this.latestMessage = atob(telemetry.Payload);
    });
   }

  ngOnInit() {
  }

  startEmulator() {
    console.log('start emulator called');
    this.telemetryService.sendCommand('Start');
  }

  stopEmulator() {
    console.log('stop emulator called');
    this.telemetryService.sendCommand('Stop');
  }

  executeNext() {
    console.log('execute next instruction called');
    this.telemetryService.sendCommand('ExecuteNext');
  }

  getCPUState() {
    console.log('get CPU state called');
    this.telemetryService.sendCommand('GetCPUState');
  }

  getEmulatorState() {
    console.log('get Emulator state called');
    this.telemetryService.sendCommand('GetEmulatorState');
  }

  getMemoryContent() {
    console.log('get Memory content called');
    this.telemetryService.sendCommand('GetMemoryContent');
  }
}
