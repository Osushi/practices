require 'bundler/setup'
require 'sinatra'

class Demo < Sinatra::Base
  get '/' do
    'Hello world!'
  end
end
