import { Component, OnInit } from '@angular/core';
import { TelemetryService } from '../services/telemetry.service';
import { Telemetry } from '../models/telemetry.model';
import { EmulatorState } from '../models/emulatorstate.model';

@Component({
  selector: 'app-emulator',
  templateUrl: './emulator.component.html',
  styleUrls: ['./emulator.component.scss']
})
export class EmulatorComponent implements OnInit {

  latestMessage: string;
  public state: EmulatorState;

  constructor(private telemetryService: TelemetryService) {
    this.state = new EmulatorState();

    telemetryService.getTelemetry().subscribe((t: string) => {
      const telemetry: Telemetry = JSON.parse(t);

      if (telemetry.Command === '"GetEmulatorState"') {
        this.state = JSON.parse(atob(telemetry.Payload));
      }
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

    setTimeout(() => {
      this.telemetryService.sendCommand('GetCPUState');
    }, 50);

    setTimeout(() => {
      this.telemetryService.sendCommand('GetEmulatorState');
    }, 100);
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
