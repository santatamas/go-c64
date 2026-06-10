import { Component, ElementRef, ViewChild } from '@angular/core';
import { MatCardModule } from '@angular/material/card';
import { TelemetryService } from '../services/telemetry.service';
import { Telemetry } from '../models/telemetry.model';
import { CPUState } from '../models/cpustate.model';

@Component({
  selector: 'app-memory',
  standalone: true,
  imports: [MatCardModule],
  template: `
    <mat-card>
      <mat-card-header><mat-card-title>Memory</mat-card-title><mat-card-subtitle>Program counter window</mat-card-subtitle></mat-card-header>
      <mat-card-content #dataContainer class="monospace" style="font-size:12px;overflow-x:auto;white-space:nowrap;"></mat-card-content>
    </mat-card>
  `
})
export class MemoryComponent {
  @ViewChild('dataContainer') dataContainer!: ElementRef;
  private cpuState: CPUState = { A: 0, Y: 0, X: 0, S: 0, P: 0, PC: 0, SP: 0, SP_LOW: 0, SP_HIGH: 0, InstructionType: 0, InstructionName: '' };

  constructor(private telemetryService: TelemetryService) {
    telemetryService.getTelemetry().subscribe((t: string) => {
      const telemetry: Telemetry = JSON.parse(t);

      if (telemetry.Command === 'GetCPUState') {
        this.cpuState = JSON.parse(atob(telemetry.Payload));
      }

      if (telemetry.Command === 'GetMemoryContent') {
        const byteCharacters = atob(telemetry.Payload);
        const byteNumbers = new Array(byteCharacters.length);
        for (let i = 0; i < byteCharacters.length; i++) {
          byteNumbers[i] = byteCharacters.charCodeAt(i);
        }
        const byteArray = new Uint8Array(byteNumbers);
        let hexResult = '';
        const windowSize = 100;
        let windowStart = this.cpuState.PC - windowSize;
        while (windowStart % 16 !== 0) windowStart--;
        let windowEnd = this.cpuState.PC + windowSize;
        while (windowEnd % 16 !== 15) windowEnd++;

        for (let cnt = 0; cnt < byteNumbers.length; cnt++) {
          if (cnt >= windowStart && cnt <= windowEnd) {
            if (cnt === this.cpuState.PC) hexResult += '<span class="current-pc-highlight">';
            hexResult += (cnt % 16 ? ' ' : '<br/>' + ('0000000' + cnt.toString(16)).slice(-8) + ' | ') + ('0' + byteArray[cnt].toString(16)).slice(-2);
            if (cnt === this.cpuState.PC) hexResult += '</span>';
          }
        }
        if (this.dataContainer) {
          this.dataContainer.nativeElement.innerHTML = hexResult;
        }
      }
    });
  }
}
