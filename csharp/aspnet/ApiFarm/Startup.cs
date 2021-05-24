#pragma warning disable
using ApiFarm.Models.Impl;
using ApiFarm.Repositories;
using ApiFarm.Services;
using ApiFarm.Services.Impl;
using ApiFarm.Utils;
using ApiFarm.Utils.Impl;
using Microsoft.AspNetCore.Builder;
using Microsoft.AspNetCore.Hosting;
using Microsoft.AspNetCore.Mvc;
using Microsoft.Extensions.Configuration;
using Microsoft.Extensions.DependencyInjection;

namespace ApiFarm.Scaffolding
{
    public class Startup
    {
        public Startup(IConfiguration configuration)
        {
            this.Configuration = configuration;
        }

        public IConfiguration Configuration { get; }

        public void ConfigureServices(IServiceCollection services)
        {
            services.AddMvc().SetCompatibilityVersion(CompatibilityVersion.Version_2_1);

            services.AddTransient<IRepository<VideoGame>, InMemory<VideoGame>>();

            services.AddTransient<IService<VideoGame>, VideoGameService>();

            services.AddTransient<IQueryFactory, QueryFactory>();
        }

        public void Configure(IApplicationBuilder app, IHostingEnvironment env)
        {
            if (env.IsDevelopment())
            {
                app.UseDeveloperExceptionPage();
            }

            app.UseMvc();
        }
    }
}
