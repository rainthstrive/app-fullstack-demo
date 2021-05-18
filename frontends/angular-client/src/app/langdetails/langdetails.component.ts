import { Component, OnInit } from '@angular/core';
import { ActivatedRoute } from '@angular/router';
import { Location } from '@angular/common';
import { ProglangsService } from '../proglangs.service';
import { proglangs } from '../PROGLANGS';

@Component({
  selector: 'app-langdetails',
  templateUrl: './langdetails.component.html',
  styleUrls: ['./langdetails.component.scss']
})
export class LangdetailsComponent implements OnInit {
  lang: proglangs | undefined;
  // TODO: GENERAR REACTIVE FORM
  constructor(
    private route: ActivatedRoute,
    private service: ProglangsService,
    private location: Location
  ) { }

  ngOnInit(): void {
    this.getLang();
  }

  getLang(): void {
    const id = parseInt(this.route.snapshot.paramMap.get('id')!, 10);
    this.service.getLang(id)
      .subscribe(res => this.lang = res);
  }

  goBack(): void {
    this.location.back();
  }

  save(): void {
    // TODO: USAR LA LOGICA DE ADD PARA GUARDAR CAMBIOS
    if (this.lang) {
      console.log(this.lang)
      //this.service.updateLang(this.lang)
      //  .subscribe(() => this.goBack());
    }
  }

}
