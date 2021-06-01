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
using Microsoft.AspNetCore.Mvc.ApplicationParts;
using Microsoft.Extensions.Configuration;
using Microsoft.Extensions.DependencyInjection;
using Microsoft.FeatureManagement;

namespace ApiFarm.Scaffolding
{
    public class Startup
    {
        public Startup(IConfiguration configuration, IHostingEnvironment env)
        {
            this.Configuration = configuration;
            this.HostingEnvironment = env;
        }

        public IConfiguration Configuration { get; }
        public IHostingEnvironment HostingEnvironment { get; }


        public void ConfigureServices(IServiceCollection services)
        {
            services.AddMvc().SetCompatibilityVersion(CompatibilityVersion.Version_2_1);

            services.AddFeatureManagement();

            services.AddSingleton<IRepository<VideoGame>, InMemory<VideoGame>>();

            services.AddSingleton<IService<VideoGame>, VideoGameService>();

            services.AddSingleton<IQueryFactory, QueryFactory>();

            services.AddSingleton<ICloner<VideoGame>, VideoGameCloner>();

            services.AddSingleton<IDataLoader<VideoGame>, VideoGameDataLoader>();
        }

        public void Configure(IApplicationBuilder app)
        {
            if (HostingEnvironment.IsEnvironment("dev"))
            {
                app.UseDeveloperExceptionPage();
            }

            app.UseMvc();
        }
    }
}
