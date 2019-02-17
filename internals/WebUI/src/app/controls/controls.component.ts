import { Component, OnInit } from '@angular/core';
import { EmulatorState } from '../models/emulatorstate.model';
import { TelemetryService } from '../services/telemetry.service';
import { Telemetry } from '../models/telemetry.model';

@Component({
  selector: 'app-controls',
  templateUrl: './controls.component.html',
  styleUrls: ['./controls.component.scss']
})
export class ControlsComponent implements OnInit {

  latestMessage: string;
  public state: EmulatorState;

  constructor(private telemetryService: TelemetryService) {
    this.state = new EmulatorState();

    telemetryService.getTelemetry().subscribe((t: string) => {
      const telemetry: Telemetry = JSON.parse(t);

      if (telemetry.Command === 'GetEmulatorState') {
        this.state = JSON.parse(atob(telemetry.Payload));
      }
    });
   }

  ngOnInit() {
  }

  setBreakpoint(value: string) {
    const telemetryRequest = new Telemetry();
    telemetryRequest.Command = 'SetBreakpoint';
    // convert from hex to dec
    const dec = parseInt(value, 16);
    // set the decimal number as a parameter on the command
    telemetryRequest.Parameter = dec.toString();
    this.telemetryService.sendCommand(telemetryRequest);
  }

  startEmulator() {
    console.log('start emulator called');
    this.telemetryService.sendStringCommand('Start');

    this.refreshAll();
  }

  stopEmulator() {
    console.log('stop emulator called');
    this.telemetryService.sendStringCommand('Stop');

    this.refreshAll();
  }

  executeNext() {
    console.log('execute next instruction called');
    this.telemetryService.sendStringCommand('ExecuteNext');

    this.refreshAll();
  }

  getCPUState() {
    console.log('get CPU state called');
    this.telemetryService.sendStringCommand('GetCPUState');
  }

  getEmulatorState() {
    console.log('get Emulator state called');
    this.telemetryService.sendStringCommand('GetEmulatorState');
  }

  getMemoryContent() {
    console.log('get Memory content called');
    this.telemetryService.sendStringCommand('GetMemoryContent');
  }

  refreshAll() {
    setTimeout(() => {
      this.telemetryService.sendStringCommand('GetCPUState');
    }, 50);

    setTimeout(() => {
      this.telemetryService.sendStringCommand('GetEmulatorState');
    }, 100);

    setTimeout(() => {
      this.telemetryService.sendStringCommand('GetMemoryContent');
    }, 150);
  }

}
