namespace TradersWrapper;


public interface ISpaceTradersClient
{
    Task<string> GetStatusAsync();
    Task<string> GetUserAsync();

}


public class Client
{
    public Client()
    {

    }
}
