using System;
using System.Linq;
using System.Text.Json;
using Microsoft.AspNetCore.Builder;
using Microsoft.AspNetCore.Hosting;
using Microsoft.AspNetCore.Http;
using Microsoft.Extensions.DependencyInjection;
using Microsoft.Extensions.Hosting;

namespace api_user
{
  public class Startup
  {
    public void ConfigureServices(IServiceCollection services)
    {
    }

    string[] words = { "lorem", "ipsum", "dolor", "sit", "amet", "consectetuer", "adipiscing", "elit", "sed", "diam", "nonummy", "nibh", "euismod", "tincidunt", "ut", "laoreet", "dolore", "magna", "aliquam", "erat" };
    Random random = new Random();

    public void Configure(IApplicationBuilder app, IWebHostEnvironment env)
    {
      if (env.IsDevelopment())
      {
        app.UseDeveloperExceptionPage();
      }

      app.UseRouting();

      app.UseEndpoints(endpoints =>
      {
        endpoints.MapGet("/", async context =>
        {
          await context.Response.WriteAsync($"Hello WorldWide Software Architecture Summit 2021 ({ Environment.MachineName })!");
        });

        endpoints.MapPost("/get-user-list", async context =>
        {
          var list = Enumerable.Range(1, random.Next(5, 15)).Select(index => new
          {
            BirthDate = DateTime.Now.AddDays(index),
            Salary = ((random.NextDouble() * 100000) + 50000).ToString("0.##"),
            Name = words[random.Next(words.Length)] + " " + words[random.Next(words.Length)],
            ProfilePictureUrl = "https://i.pravatar.cc/50?" + random.Next(1, 1000)
          });

          await JsonSerializer.SerializeAsync(context.Response.Body, list);
        });
      });
    }
  }
}
