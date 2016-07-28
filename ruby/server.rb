require 'sinatra'
require 'httparty'
require 'logger'

class MyApp < Sinatra::Base
  configure :production, :development do
    set :logging, true
    set :dump_errors, true
  end

  get '/' do
    #HTTParty.get('www.whatever.com')
    #logger.info 'httparty and logging works'

    'Hello world'
  end
end
