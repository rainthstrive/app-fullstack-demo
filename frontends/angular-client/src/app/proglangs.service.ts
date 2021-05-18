import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import { proglangs } from './PROGLANGS';

@Injectable({
  providedIn: 'root'
})
export class ProglangsService {
  goBaseUrl = 'http://localhost:1323/prog_langs';
  expressBaseUrl = 'http://localhost:8080/prog_langs';
  constructor(private http: HttpClient){
  }

  getLangs(): Observable<proglangs[]> {
    return this.http.get<proglangs[]>(this.goBaseUrl);
  }

  getLang(id: number): Observable<proglangs> {
    const url = `${this.goBaseUrl}/${id}`;
    return this.http.get<proglangs>(url);
  }

  addLang(lang: proglangs): Observable<proglangs> {
    const formData = new FormData();
    //console.log('SERVICE: ',lang);
    formData.append('name', lang.name);
    formData.append('auth', lang.auth);
    formData.append('comp', lang.comp);
    formData.append('rel_date', String(lang.rel_date));
    return this.http.post<any>(this.goBaseUrl, formData);
  }

  deleteLang(id: number): Observable<proglangs> {
    const url = `${this.goBaseUrl}/${id}`;
    return this.http.delete<proglangs>(url);
  }

  updateLang(lang: proglangs): Observable<any> {
    const url = `${this.goBaseUrl}/${lang.id}`;
    return this.http.put(this.goBaseUrl, lang);
  }

}
