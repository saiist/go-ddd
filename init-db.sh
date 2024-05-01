#!/bin/bash

psql -h localhost -U postgres -d postgres -f ./init-db.sql
