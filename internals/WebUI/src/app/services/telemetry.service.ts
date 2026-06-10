import { Injectable } from '@angular/core';
import { Subject } from 'rxjs';
import { WebsocketService } from './websocket.service';
import { Telemetry } from '../models/telemetry.model';

@Injectable({ providedIn: 'root' })
export class TelemetryService {
  private telemetry$: Subject<string>;

  constructor(private webSocket: WebsocketService) {
    this.telemetry$ = this.webSocket.connect('ws://localhost:8080/ws');
  }

  getTelemetry(): Subject<string> {
    return this.telemetry$;
  }

  sendStringCommand(command: string): void {
    const request: Telemetry = { Command: command };
    this.sendCommand(request);
  }

  sendCommand(request: Telemetry): void {
    this.webSocket.send(JSON.stringify(request));
  }
}
