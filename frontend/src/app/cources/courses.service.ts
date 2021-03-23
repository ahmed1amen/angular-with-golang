import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Observable, of } from 'rxjs';
import { Cource } from './cource';
import { catchError, map, tap } from 'rxjs/operators';
@Injectable({
  providedIn: 'root'
})
export class CoursesService {

  private heroesUrl = 'http://localhost:8000/courses';  // URL to web api

  httpOptions = {
    headers: new HttpHeaders({ 'Content-Type': 'application/json' })
  };
  constructor( private http: HttpClient ) {

  }
  getHeroes(): Observable<Cource[]> {
    return this.http.get<Cource[]>(this.heroesUrl);

  }
}
