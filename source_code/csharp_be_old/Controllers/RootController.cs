using System.Web.Http;

namespace CSharpOldWebApi.Controllers
{
    public class RootController : ApiController
    {
        [HttpGet]
        [Route("")]
        public object Get()
        {
            return new
            {
                framework = "ASP.NET Web API",
                framework_version = typeof(ApiController).Assembly.GetName().Version.ToString(),
                dotnet_framework = ".NET Framework 4.8"
            };
        }
    }
}
