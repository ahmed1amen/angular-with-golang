import { Component, OnInit } from '@angular/core';
import { CoursesService } from './courses.service';
import { Cource } from './cource';

@Component({
  selector: 'app-cources',
  templateUrl: './courses.component.html',
  styleUrls: ['./courses.component.css']
})
export class CoursesComponent implements OnInit {
  courses: Cource[];
  public title;

  constructor(private coursesService: CoursesService) {
    this.title = 'OOP';
    this.courses = [];
  }

  getHeroes(): void {
    this.coursesService.getHeroes()
      .subscribe(heroes => this.courses = heroes);
  }

  ngOnInit(): void {
    this.getHeroes();
  }

}
