# Monitoring - School Project
*Third Year @ Ynov Informatique Paris*
Le but du projet était de créer une application de monitoring.

## Collaborateurs
- Naomi Paulmin, Web Developer [@codebynao](https://github.com/codebynao)
- Matthieu Poulier, Web and Golang Developer, Security [@mpoulier](https://github.com/mpoulier)
- Valentin Marlier, Golang Developer, CI/CD manager [@vmarlier](https://github.com/vmarlier)

## Fonctionnalitées:
- Agrégation et affichage des logs et metrics
- Management des utilisateurs et des groups (définir qui est sudoer..)
- Management des processes (kill..)
- Utilisation d'un webshell pour manager les clients

## Application Architecture:
- [Web-App](https://github.com/codebynao/Monitoring/tree/master/Web-App): En Nodejs, affiche tous les logs/metrics, permet de gérer les groupes et les utilisateurs, lance le webshell.
- Agents: En Golang, installé sur les clients à monitorer, expose une api permettant le monitoring, le management des utilisateurs et des processus ainsi que l'utilisation du webshell.
- Orchestrator: Supposé fonctionné comme le SearchHead dans Splunk, n'a pas été développé pour le rendu final.

## Stack Technique:
<p align="left">
	<a href="https://about.gitlab.com/" target="_blank">
		<img src="https://www.vectorlogo.zone/logos/gitlab/gitlab-icon.svg" alt="gitlab" width="40" height="40"/>
	</a>
	<a href="https://www.docker.com/" target="_blank">
		<img src="https://raw.githubusercontent.com/gilbarbara/logos/master/logos/docker-icon.svg" alt="docker" width="40" height="40"/>
	</a>
	<a href="https://golang.org/" target="_blank">
		<img src="https://www.vectorlogo.zone/logos/golang/golang-icon.svg" alt="golang" width="40" height="40"/>
	</a>
	<a href="https://nodejs.org/en/" target="_blank">
		<img src="https://www.vectorlogo.zone/logos/nodejs/nodejs-horizontal.svg" alt="bash" width="40" height="40"/>
	</a>
    <a href="https://vuejs.org/" target="_blank">
		<img src="https://www.vectorlogo.zone/logos/vuejs/vuejs-icon.svg" alt="bash" width="40" height="40"/>
	</a>
</p>
