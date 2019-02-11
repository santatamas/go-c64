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
  }

  stopEmulator() {
    console.log('stop emulator called');
    this.telemetryService.sendStringCommand('Stop');
  }

  executeNext() {
    console.log('execute next instruction called');
    this.telemetryService.sendStringCommand('ExecuteNext');

    setTimeout(() => {
      this.telemetryService.sendStringCommand('GetCPUState');
    }, 50);

    setTimeout(() => {
      this.telemetryService.sendStringCommand('GetEmulatorState');
    }, 100);
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

}
