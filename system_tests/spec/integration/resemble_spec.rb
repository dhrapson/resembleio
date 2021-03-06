# Copyright (C) 2015 dhrapson

# This program is free software: you can redistribute it and/or modify
# it under the terms of the GNU General Public License as published by
# the Free Software Foundation, either version 3 of the License, or
# (at your option) any later version.

# This program is distributed in the hope that it will be useful,
# but WITHOUT ANY WARRANTY; without even the implied warranty of
# MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
# GNU General Public License for more details.

# You should have received a copy of the GNU General Public License
# along with this program.  If not, see <http://www.gnu.org/licenses/>.

require 'faraday'

describe 'The resemble Server' do


	def start_process(cmd)
		pipe_cmd_in, pipe_cmd_out = IO.pipe
		cmd_pid = Process.spawn(cmd, :out => pipe_cmd_out, :err => pipe_cmd_out)
		sleep 1
		@exitstatus = :not_done
		Thread.new do
		  Process.wait(cmd_pid)
		  @exitstatus = $?.exitstatus if $?.exited?
		end
		{
			pid: cmd_pid,
			pipe_cmd_in: pipe_cmd_in,
			pipe_cmd_out: pipe_cmd_out
		}
	end

	def terminate_process(process_details, initial_wait_seconds=0.1)
		Thread.new do
		  sleep initial_wait_seconds
		  Process.kill('KILL', process_details[:pid])
		  @exitstatus = -1
		end

		process_details[:pipe_cmd_out].close
		out = process_details[:pipe_cmd_in].read
		sleep(0.1) while @exitstatus == :not_done
		{
			output: out,
			exitstatus: @exitstatus
		}
	end

	def spawn_process(cmd,initial_wait_seconds=5)
		process_details = start_process(cmd)
		terminate_process(process_details, initial_wait_seconds)
	end

	let(:client) do
		Faraday.new(:url => 'http://localhost:9000') do |faraday|
			faraday.adapter	Faraday.default_adapter
		end
	end

	context 'when run without a config file' do

		it 'starts anyway' do
			results = spawn_process('resemble')
			# a -1 exit status indicates it had to be killed
			expect(results[:exitstatus]).to be -1
			expect(results[:output]).to include('Starting Resemble Service')
		end

		context 'when its left running' do
			attr_reader :process_details

			before(:all) do
				@process_details = start_process('resemble')
			end

			after(:all) do
				terminate_process(process_details)
			end

			it 'raises an API endpoint' do
				response = client.get('/resemble')
				expect(response.status).to be 200
			end
		end
	end

	context 'when run with a missing config file' do
		it 'errors out with a non-zero error code' do
			results = spawn_process('resemble /path/to/file/that/doesnt/exist.yml')
			expect($?).to_not be 0
			expect(results[:output]).to include('cannot be found')
		end
	end

	context 'when run with a garbage config file' do
		it 'errors out with a non-zero error code' do
			results = spawn_process('resemble spec/integration/fixtures/invalid_config.yml')
			expect($?).to_not be 0
			expect(results[:output]).to include('Invalid config: `type`')
		end
	end

	context 'when running the minimal configuration' do
		attr_reader :process_details

		before(:all) do
			@process_details = start_process('resemble spec/integration/fixtures/min_rest_resemble.yml')
		end

		after(:all) do
			terminate_process(process_details)
		end

		it 'raises an API endpoint' do
			response = client.get('/resemble')
			expect(response.status).to be 200
		end
	end

	context 'when running a very specific HTTP configuration' do
		attr_reader :process_details

		before(:all) do
			@process_details = start_process('resemble spec/integration/fixtures/http_full_resemble.yml')
		end

		after(:all) do
			terminate_process(process_details)
		end

		it 'raises an API endpoint' do
			response = client.get('/resemble')
			expect(response.status).to be 200
		end

		it 'correctly mis-matches an unconfigured URL path' do
			response = client.get('/notconfiguredurlpath?id=abc123')
			expect(response.status).to be 404
		end

		it 'correctly mis-matches an unconfigured HTTP verb' do
			response = client.put('/test?id=abc123')
			expect(response.status).to be 404
		end

		it 'correctly mis-matches an unconfigured HTTP verb' do
			response = client.put('/test?id=abc123')
			expect(response.status).to be 404
		end

		it 'correctly mis-matches a URL without required query params' do
			response = client.get('/test')
			expect(response.status).to be 404
		end

		it 'correctly matches a configured endpoint' do
			response = client.get('/test?id=abc123')
			expect(response.status).to be 200
		end

		it 'correctly uses a second matcher on the same endpoint' do
			response = client.post('/testagain')
			expect(response.status).to be 200
		end

		it 'correctly matches on a second endpoint' do
			response = client.post('/second_endpoint')
			expect(response.status).to be 200
			response = client.put('/second_endpoint')
			expect(response.status).to be 200
			response = client.get('/second_endpoint')
			expect(response.status).to be 200
		end
	end

	context 'when running a config with a non-existing host' do
		attr_reader :process_details

		before(:all) do
			@process_details = start_process('resemble spec/integration/fixtures/http_bad_host_resemble.yml')
		end

		after(:all) do
			terminate_process(process_details)
		end
		it 'correctly mis-matches the request' do
			response = client.get('/test')
			expect(response.status).to be 404
		end
	end

	context 'when running a config URL path containing a guid' do
		attr_reader :process_details

		before(:all) do
			@process_details = start_process('resemble spec/integration/fixtures/http_with_guid_resemble.yml')
		end

		after(:all) do
			terminate_process(process_details)
		end
		it 'correctly matches the request' do
			response = client.get('/test/30dd879c-ee2f-11db-8314-0800200c9a66')
			expect(response.status).to be 200
		end
	end
end
