/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package main

import "github.com/svanellewee/history-manager/cmd"

func main() {
	cmd.Execute()
}

/*
!/usr/bin/env bash
mkdir -p ~/.bash_logs
HISTORY_DB="${1:-${HOME}/.bash_logs/history.db}"
function history-init() {
	sqlite3 "${HISTORY_DB}" <<- "EOF"
	CREATE TABLE IF NOT EXISTS entry (
	  entry_id INTEGER,
	  time TIMESTAMP,
	  VALUE VARCHAR
	);
	EOF
}

function history-add() {
	local value="${*}"
	local num="$(echo "${value}" | grep -P "[ \t]+\K[0-9]+" -o)"
	local therest="$(echo "${value}" | cut -d' ' -f3- | tr "'" "\'")"
	sqlite3 "${HISTORY_DB}" <<- EOF
	INSERT INTO entry(entry_id, time, value) VALUES (${num}, DATETIME(), '${therest}');
	EOF
}

function history-list() {
	sqlite3 "${HISTORY_DB}" <<- EOF
	.mode csv
	.headers on
	SELECT entry_id, time, value FROM entry;
	EOF
}

function history-export() {
	sqlite3 "${HISTORY_DB}" <<- EOF
	SELECT value FROM entry ORDER BY rowid ASC
	EOF
}
*/
