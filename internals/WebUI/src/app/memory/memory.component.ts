import { Component, OnInit, ElementRef, ViewChild } from '@angular/core';
import { TelemetryService } from '../services/telemetry.service';
import { Telemetry } from '../models/telemetry.model';
import { CPUState } from '../models/cpustate.model';

@Component({
  selector: 'app-memory',
  templateUrl: './memory.component.html',
  styleUrls: ['./memory.component.scss']
})
export class MemoryComponent implements OnInit {

  public memoryContent = 'blank';
  public cpuState: CPUState;

  @ViewChild('dataContainer') dataContainer: ElementRef;

  loadData(data) {
      this.dataContainer.nativeElement.innerHTML = data;
  }

  constructor(private telemetryService: TelemetryService) {

    this.cpuState = new CPUState();

    telemetryService.getTelemetry().subscribe((t: string) => {
      const telemetry: Telemetry = JSON.parse(t);

      if (telemetry.Command === 'GetCPUState') {
        this.cpuState = JSON.parse(atob(telemetry.Payload));
      }

      if (telemetry.Command === 'GetMemoryContent') {

        let byteCharacters = atob(telemetry.Payload);
        let byteNumbers = new Array(byteCharacters.length);

        for (let i = 0; i < byteCharacters.length; i++) {
          byteNumbers[i] = byteCharacters.charCodeAt(i);
        }

        let byteArray = new Uint8Array(byteNumbers);

        //console.log(byteArray);
        let hexResult = '';
        let cnt = 0;
        let windowSize = 100;
        let windowStart = 0;
        let windowEnd = 0;

        windowStart = this.cpuState.PC - windowSize;
        while (windowStart % 16 !== 0) {
          windowStart = windowStart - 1;
        }
        windowEnd = this.cpuState.PC + windowSize;
        while (windowEnd % 16 !== 15) {
          windowEnd = windowEnd + 1;
        }

        for (let byte of byteNumbers) {

          // only print the window
          if (cnt >= windowStart && cnt <= windowEnd) {

            // highlight the current PC instruction
            if (cnt === this.cpuState.PC) {
              hexResult += '<span style="background-color: yellow;color: red;" class="current-pc-highlight">';
            }

            // tslint:disable-next-line:max-line-length
            hexResult += (cnt % 16 ? ' ' : '<br/>' + (1e7 + (cnt).toString(16)).slice(-8) + ' | ') + (1e7 + byteArray[cnt].toString(16)).slice(-2);

            if (cnt === this.cpuState.PC) {
              hexResult += '</span>';
            }
          }
          cnt++;
        }

        this.dataContainer.nativeElement.innerHTML = hexResult;
      }
    });
  }

  ngOnInit() {
  }

}
