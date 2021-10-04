import { Component, OnInit } from '@angular/core';
import { proglangs } from '../PROGLANGS';
import { ProglangsService } from '../proglangs.service';
import { FormGroup, FormControl } from '@angular/forms';

@Component({
  selector: 'app-langlist',
  templateUrl: './langlist.component.html',
  styleUrls: ['./langlist.component.scss']
})
export class LanglistComponent implements OnInit {
  form = new FormGroup({
    name: new FormControl(''),
    auth: new FormControl(''),
    comp: new FormControl(''),
    rel_date: new FormControl(''),
  });
  prolangs: proglangs[] = [];
  constructor(private service: ProglangsService) {

  }
  ngOnInit(){
    this.getLangs();
  }

  // Función para Retornar arreglo de registros
  getLangs(){
    this.service.getLangs()
    .subscribe(langs => {this.prolangs = langs; console.log(this.prolangs)})
  }

  add(): void {
    if (!this.form.value) { return; }
    const input: proglangs = {
      id: 0,
      name: this.form.value.name,
      comp: this.form.value.comp,
      auth: this.form.value.auth,
      rel_date: this.getEpoch(this.form.value.rel_date)
    };
    //console.log('INPUTS: ', input);
    this.service.addLang(input)
      .subscribe(
        (res) => {this.prolangs.push(res); console.log(res)},
        (err) => console.log(err)
    );
  }

  delete(id: number): void {
    this.prolangs = this.prolangs.filter(el => el.id !== id);
    this.service.deleteLang(id).subscribe();
  }

  getEpoch(date: string): number {
    const timestamp = new Date(date).getTime();
    return Math.floor( timestamp / 1000);
  }
}
