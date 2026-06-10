import { Component } from '@angular/core';
import { CommonModule } from '@angular/common';
import { MatButtonModule } from '@angular/material/button';
import { MatIconModule } from '@angular/material/icon';
import { TelemetryService } from '../services/telemetry.service';
import { Telemetry } from '../models/telemetry.model';
import { EmulatorState } from '../models/emulatorstate.model';

@Component({
  selector: 'app-controls',
  standalone: true,
  imports: [CommonModule, MatButtonModule, MatIconModule],
  template: `
    <button *ngIf="state.PauseFlag" (click)="startEmulator()" title="Start" mat-icon-button><mat-icon>play_arrow</mat-icon></button>
    <button *ngIf="!state.PauseFlag" (click)="stopEmulator()" title="Pause" mat-icon-button><mat-icon>pause</mat-icon></button>
    <button (click)="executeNext()" title="Step" mat-icon-button><mat-icon>skip_next</mat-icon></button>
    <button (click)="getCPUState()" title="CPU State" mat-icon-button><mat-icon>memory</mat-icon></button>
    <button (click)="getEmulatorState()" title="Emulator State" mat-icon-button><mat-icon>computer</mat-icon></button>
    <button (click)="getCIAState()" title="CIA State" mat-icon-button><mat-icon>keyboard</mat-icon></button>
    <button (click)="getMemoryContent()" title="Memory" mat-icon-button><mat-icon>sd_card</mat-icon></button>
    <input #box placeholder="Breakpoint (hex)" style="margin-left:8px;width:100px;">
    <button (click)="setBreakpoint(box.value)" mat-button>Set BP</button>
  `
})
export class ControlsComponent {
  state: EmulatorState = { Delay: '0', CycleCount: '0', PauseFlag: true, Debug: false, Test: false, BreakAddr: 0 };

  constructor(private telemetryService: TelemetryService) {
    telemetryService.getTelemetry().subscribe((t: string) => {
      const telemetry: Telemetry = JSON.parse(t);
      if (telemetry.Command === 'GetEmulatorState') {
        this.state = JSON.parse(atob(telemetry.Payload));
      }
    });
  }

  setBreakpoint(value: string): void {
    const request: Telemetry = { Command: 'SetBreakpoint', Parameter: parseInt(value, 16).toString() };
    this.telemetryService.sendCommand(request);
  }

  startEmulator(): void { this.telemetryService.sendStringCommand('Start'); this.refreshAll(); }
  stopEmulator(): void { this.telemetryService.sendStringCommand('Stop'); this.refreshAll(); }
  executeNext(): void { this.telemetryService.sendStringCommand('ExecuteNext'); this.refreshAll(); }
  getCPUState(): void { this.telemetryService.sendStringCommand('GetCPUState'); }
  getCIAState(): void { this.telemetryService.sendStringCommand('GetCIAState'); }
  getEmulatorState(): void { this.telemetryService.sendStringCommand('GetEmulatorState'); }
  getMemoryContent(): void { this.telemetryService.sendStringCommand('GetMemoryContent'); }

  private refreshAll(): void {
    setTimeout(() => this.telemetryService.sendStringCommand('GetCPUState'), 50);
    setTimeout(() => this.telemetryService.sendStringCommand('GetCIAState'), 75);
    setTimeout(() => this.telemetryService.sendStringCommand('GetEmulatorState'), 100);
    setTimeout(() => this.telemetryService.sendStringCommand('GetMemoryContent'), 125);
  }
}
