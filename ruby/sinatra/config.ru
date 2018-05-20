require 'rubygems'
require 'bundler/setup'
require 'rack'
require 'sinatra'
if development?
  $stdout.sync = true
  require 'sinatra/reloader'
end
require 'json'
$:.unshift File.dirname(__FILE__)
require 'helpers/helper'
require 'controllers/main'

run Sinatra::Application
