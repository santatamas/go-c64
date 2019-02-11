import { Component, OnInit, ElementRef, ViewChild } from '@angular/core';
import { TelemetryService } from '../services/telemetry.service';
import { Telemetry } from '../models/telemetry.model';

@Component({
  selector: 'app-memory',
  templateUrl: './memory.component.html',
  styleUrls: ['./memory.component.scss']
})
export class MemoryComponent implements OnInit {

  public memoryContent = 'blank';

  @ViewChild('dataContainer') dataContainer: ElementRef;

  loadData(data) {
      this.dataContainer.nativeElement.innerHTML = data;
  }

  constructor(private telemetryService: TelemetryService) {

    telemetryService.getTelemetry().subscribe((t: string) => {
      const telemetry: Telemetry = JSON.parse(t);

      if (telemetry.Command === '"GetMemoryContent"') {

        let byteCharacters = atob(telemetry.Payload);

        let byteNumbers = new Array(byteCharacters.length);

        for (let i = 0; i < byteCharacters.length; i++) {
          byteNumbers[i] = byteCharacters.charCodeAt(i);
        }

        let byteArray = new Uint8Array(byteNumbers);


        console.log(byteArray);
        let hexResult = '';
        let cnt = 0;

        for (let byte of byteNumbers) {
          // tslint:disable-next-line:max-line-length
          hexResult += (cnt % 16 ? ' ' : '<br/>' + (1e7 + (cnt).toString(16)).slice(-8) + ' | ') + (1e7 + byteArray[cnt].toString(16)).slice(-2);
          cnt++;
        }

        this.dataContainer.nativeElement.innerHTML = hexResult;
      }
    });
  }

  ngOnInit() {
  }

}
